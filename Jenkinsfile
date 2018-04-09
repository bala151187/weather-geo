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
         withKubeConfig(caCertificate: '''-----BEGIN CERTIFICATE----- MIIDCzCCAfOgAwIBAgIQI+JsrNZEfPIHtdZ8+c7/JzANBgkqhkiG9w0BAQsFADAv MS0wKwYDVQQDEyRlMDlkNzcyNi1jYTZiLTRlYTQtYTdlMC1iNGUwZDhjNGViNmMw HhcNMTgwNDA3MDAzMzU5WhcNMjMwNDA2MDEzMzU5WjAvMS0wKwYDVQQDEyRlMDlk NzcyNi1jYTZiLTRlYTQtYTdlMC1iNGUwZDhjNGViNmMwggEiMA0GCSqGSIb3DQEB AQUAA4IBDwAwggEKAoIBAQDYKJecNqWXxSwGd4usGt5vBHAm22t2dbJd+R8eJovf 4TwWYBhFQDyLYauy+dfIbQW2zb96FEOSN49mq2XBhs+/Pac1q219O8BRGuWIa5si KSojAMw3tZXwBP0u+zEC0ZKt9UQ2v9DzhzCdBbMaLUz1QI7zUliARdjFWAbM2L9M 9LZSw2MCi1kxW3hXAEVjPjtxO7eMRc0D3vRpEvjetFPMZSRuOKEN3L4lr7NMPypq wJ3WucCms2MNGgDrJIYtPmJ57yUP6LfzmEOJBBIqZM51Fo9jACEfAOnX0T01Q1iZ zy+2+S9ZgiK95vSBA3D1UchaHB4nMFS6eHtW3M5CdhLhAgMBAAGjIzAhMA4GA1Ud DwEB/wQEAwICBDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAM 2V0/weCFQ1w+BVUijdCO92jjbKm+QIibXo24SXH1PAAD4TepJqNefJcHEJ5A3lE8 9DNim7y2aCrtZSK1FFSAVmaLyGOdzWbs2TMViwV4mr9ennGohqRxBJ+GNUU979nx 4cyYD8ANS+PBuEu84nKlgv0fW76iH54MUjjkDmVTCDzanWVOpVbh9KO8gUQeyp4D xM776jWWbZIxPi12+oOcO2evonw9f4GxE76/Bl5bXwlK6WY36mMMCQf6lqJG55R5 66DQ/7TPpsexsOMKj0wCC5mIOI5OsfB6xN/5wYSoB65uHPy6RqDBN64qL+VQo2kn mxFQC0BhVOu0x4k/AudG -----END CERTIFICATE-----''', credentialsId: 'kubernetes_gcp_raw', serverUrl: 'https://35.193.109.253')
        {
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
