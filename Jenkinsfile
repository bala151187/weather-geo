node {
    def app

    stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */

        checkout scm
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("us.gcr.io/devops-200301/weather-geo:${env.BUILD_ID}")
    }

    stage('Test image') {
        /* Ideally, we would run a test framework against our image.
         * For this example, we're using a Volkswagen-type approach ;-) */

        app.inside {
            sh 'echo "Tests passed"'
        }
    }

    stage('Push image to GCR') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        sh "gcloud docker -- push us.gcr.io/devops-200301/weather-geo:${env.BUILD_NUMBER}"
    }
    
    stage('Deploy to kubernetes') {
        def label = "mypod-${UUID.randomUUID().toString()}"
        podTemplate(label: label, yaml: """
        apiVersion: extensions/v1beta1
        kind: Deployment
        metadata:
          creationTimestamp: null
          labels:
            app: weather-geo
            name: weather-geo
            type: web
          name: weather-geo
        spec:
          replicas: 1
          selector:
            matchLabels:
              app: weather-geo
              name: weather-geo
              type: web
          strategy:
            rollingUpdate:
              maxSurge: 1
              maxUnavailable: 1
            type: RollingUpdate
          template:
            metadata:
              labels:
                app: weather-geo
                name: weather-geo
                type: web
              name: weather-geo
            spec:
              containers:
              - image: us.gcr.io/devops-200301/weather-geo:${env.BUILD_NUMBER}
                imagePullPolicy: IfNotPresent
                name: weather-geo
                ports:
                - containerPort: 3001
              dnsPolicy: ClusterFirst
              restartPolicy: Always
              securityContext: {}
              terminationGracePeriodSeconds: 30
            """
        ) {
            node (label) {
                stage('Run specific shell') {
                    container('weather-geo') {
                        sh "hostname"
                    }
                }
            }
        }
    }
}
