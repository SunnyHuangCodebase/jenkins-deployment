FROM jenkins/jenkins

# Expose ports
EXPOSE 8080
EXPOSE 50000

# Volume for Jenkins home directory
VOLUME /var/jenkins_home

# Set container name
container_name jenkins

# Start Jenkins
CMD ["jenkins", "start"]

#FROM jenkins/jenkins
#
#EXPOSE 8080
#EXPOSE 50000
#
#VOLUME /var/jenkins_home
#
#CMD ["jenkins", "/sbin/tini", "--", "/usr/local/bin/jenkins.sh"]
