package main

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	apps "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	core "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	meta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		appName := "jenkins"
		appLabels := pulumi.StringMap{
			"app": pulumi.String(appName),
		}

		deployment, err := apps.NewDeployment(ctx, "jenkins-deployment", &apps.DeploymentArgs{
			Spec: apps.DeploymentSpecArgs{
				Selector: &meta.LabelSelectorArgs{
					MatchLabels: appLabels,
				},
				Replicas: pulumi.Int(1),
				Template: &core.PodTemplateSpecArgs{
					Metadata: &meta.ObjectMetaArgs{
						Labels: appLabels,
					},
					Spec: &core.PodSpecArgs{
						Containers: core.ContainerArray{
							core.ContainerArgs{
								Name:  pulumi.String(appName),
								Image: pulumi.String("jenkins/jenkins"),
								Ports: core.ContainerPortArray{
									core.ContainerPortArgs{
										ContainerPort: pulumi.Int(8080),
										Name:          pulumi.String("http"),
									},
								},
							},
						},
						Volumes: core.VolumeArray{
							core.VolumeArgs{
								Name: pulumi.String("jenkins-data"),
								EmptyDir: &core.EmptyDirVolumeSourceArgs{
									Medium: pulumi.String("Memory"),
								},
							},
						},
					},
				},
			},
		})

		if err != nil {
			return err
		}

		frontend, err := core.NewService(ctx, appName, &core.ServiceArgs{
			Metadata: &meta.ObjectMetaArgs{
				Labels: appLabels,
			},
			Spec: &core.ServiceSpecArgs{
				Type: pulumi.String("ClusterIP"),
				Ports: &core.ServicePortArray{
					core.ServicePortArgs{
						Port:       pulumi.Int(8000),
						TargetPort: pulumi.Int(8080),
						Protocol:   pulumi.String("TCP"),
					},
				},
				Selector: appLabels,
			},
		})

		ctx.Export("name", deployment.Metadata.Name())
		ctx.Export("frontendIp", frontend.Spec.ClusterIP())

		authorization, err := local.NewCommand(ctx, "authorization", &local.CommandArgs{
			Create: pulumi.String("kubectl exec -i `kubectl get pod -o name | grep jenkins` -- sh -c 'cat /var/jenkins_home/secrets/initialAdminPassword'"),
		}, pulumi.DependsOn([]pulumi.Resource{deployment, frontend}))

		if err != nil {
			return err
		}

		ctx.Export("jenkinsPassword", authorization.Stdout)

		return nil
	})
}
