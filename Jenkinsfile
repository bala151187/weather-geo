node {
    def app

    stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */

        checkout scm
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("us.gcr.io/devops-200301/weather-geo:latest")
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
        sh "gcloud docker -- push us.gcr.io/devops-200301/weather-geo:latest"
    }
     stage('Check Connection to k8s') {
         withKubeConfig(caCertificate: 'LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURDekNDQWZPZ0F3SUJBZ0lRSStKc3JOWkVmUElIdGRaOCtjNy9KekFOQmdrcWhraUc5dzBCQVFzRkFEQXYKTVMwd0t3WURWUVFERXlSbE1EbGtOemN5TmkxallUWmlMVFJsWVRRdFlUZGxNQzFpTkdVd1pEaGpOR1ZpTm1NdwpIaGNOTVRnd05EQTNNREF6TXpVNVdoY05Nak13TkRBMk1ERXpNelU1V2pBdk1TMHdLd1lEVlFRREV5UmxNRGxrCk56Y3lOaTFqWVRaaUxUUmxZVFF0WVRkbE1DMWlOR1V3WkRoak5HVmlObU13Z2dFaU1BMEdDU3FHU0liM0RRRUIKQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURZS0plY05xV1h4U3dHZDR1c0d0NXZCSEFtMjJ0MmRiSmQrUjhlSm92Zgo0VHdXWUJoRlFEeUxZYXV5K2RmSWJRVzJ6Yjk2RkVPU040OW1xMlhCaHMrL1BhYzFxMjE5TzhCUkd1V0lhNXNpCktTb2pBTXczdFpYd0JQMHUrekVDMFpLdDlVUTJ2OUR6aHpDZEJiTWFMVXoxUUk3elVsaUFSZGpGV0FiTTJMOU0KOUxaU3cyTUNpMWt4VzNoWEFFVmpQanR4TzdlTVJjMEQzdlJwRXZqZXRGUE1aU1J1T0tFTjNMNGxyN05NUHlwcQp3SjNXdWNDbXMyTU5HZ0RySklZdFBtSjU3eVVQNkxmem1FT0pCQklxWk01MUZvOWpBQ0VmQU9uWDBUMDFRMWlaCnp5KzIrUzlaZ2lLOTV2U0JBM0QxVWNoYUhCNG5NRlM2ZUh0VzNNNUNkaExoQWdNQkFBR2pJekFoTUE0R0ExVWQKRHdFQi93UUVBd0lDQkRBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFBTQoyVjAvd2VDRlExdytCVlVpamRDTzkyampiS20rUUlpYlhvMjRTWEgxUEFBRDRUZXBKcU5lZkpjSEVKNUEzbEU4CjlETmltN3kyYUNydFpTSzFGRlNBVm1hTHlHT2R6V2JzMlRNVml3VjRtcjllbm5Hb2hxUnhCSitHTlVVOTc5bngKNGN5WUQ4QU5TK1BCdUV1ODRuS2xndjBmVzc2aUg1NE1VamprRG1WVENEemFuV1ZPcFZiaDlLTzhnVVFleXA0RAp4TTc3NmpXV2JaSXhQaTEyK29PY08yZXZvbnc5ZjRHeEU3Ni9CbDViWHdsSzZXWTM2bU1NQ1FmNmxxSkc1NVI1CjY2RFEvN1RQcHNleHNPTUtqMHdDQzVtSU9JNU9zZkI2eE4vNXdZU29CNjV1SFB5NlJxREJONjRxTCtWUW8ya24KbXhGUUMwQmhWT3UweDRrL0F1ZEcKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=', credentialsId: 'kubernetes_GCP', serverUrl: 'https://35.193.109.253') {
    sh 'kubectl get pods'
}

    }  
    stage('Deploy to kubernetes') {
        kubernetesDeploy(kubeconfigId: 'kubernetes_GCP',             
                 configs: '**/*.yaml',
                 enableConfigSubstitution: true
        )
    }
}
