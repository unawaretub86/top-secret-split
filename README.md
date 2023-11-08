# TOP-SECRET
Go rest para obtener la ubicación de una nave y el mensaje que esta emite

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Tabla de contenidos</summary>
  <ol>
    <li>
      <a href="#sobre-el-proyecto">Sobre el proyecto</a>
      <ul>
        <li><a href="#hecho-con">Hecho con</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contacto">Contacto</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## Sobre el proyecto

Este proyecto ha sido desarrollado en el contexto de la prueba técnica Fuego de Quasar. En esta prueba, se presenta un servicio FaaS (Function as a Service) que utiliza una función Lambda serverless implementada en Golang. Esta función se encarga de guardar la información de los satélites en una base de datos DynamoDB y se comunica mediante HTTP con el servicio topsecret.

El propósito de este servicio es almacenar la información de los satélites que sean enviados a través de HTTP. Una vez se haya recibido la cantidad necesaria de satélites para llevar a cabo el cálculo, se podrá utilizar el método GET para comunicarse con el servicio topsecret. Esto permite reutilizar la funcionalidad del servicio topsecret para calcular la posición y descifrar el mensaje de manera efectiva.

Este servicio utiliza una arquitectura hexagonal diseñada para preservar la integridad de la lógica de negocio y permitir la incorporación de nuevas implementaciones a través de adaptadores, sin necesidad de modificar la lógica existente. Esta estructura proporciona flexibilidad para extender y adaptar el servicio a futuras necesidades, garantizando la coherencia de la aplicación.

En este servicio, hemos incorporado flujos de trabajo de GitHub Actions para asegurar la alta calidad de integración de Git Flow a través de pull requests. Estos flujos de trabajo incluyen acciones como construcción (build), cobertura superior al 80% (coverage) y verificación de código (lint). Esta implementación garantiza una integración fluida y una revisión exhaustiva del código, lo que nos permite mantener altos estándares de calidad en nuestro proyecto.

![actions](docs/actions.png)

<p align="right">(<a href="#readme-top">Volver al inicio</a>)</p>

### Hecho con

* [![Go][Go.dev]][Go-url]
* [![Aws][aws.com]][Aws-url]
* [![Dynamo][dynamo.com]][Dynamo-url]
* [![Github][github.com]][Github-url]
* [![Github Actions][github-actions.com]][GithubActions-url]

<p align="right">(<a href="#readme-top">Volver al inicio</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

* [Top Secret split POST](docs/topsecretsplitPost.md) : `POST /topsecret_split/{satellite_name}`
* [Top Secret split GET](docs/topsecretsplitGet.md) : `GET /topsecret_split`

<p align="right">(<a href="#readme-top">Volver al inicio</a>)</p>

<!-- CONTACT -->
## Contacto

Esteban Gomez - [LinkedIn](https://www.linkedin.com/public-profile/settings?lipi=urn%3Ali%3Apage%3Ad_flagship3_profile_self_edit_contact-info%3BSnAu1%2F2AQQiK8pzcnqrLUA%3D%3D) - goal.esteban10@gmail.com

Project Link: [Github repo](https://github.com/unawaretub86/top-secret)

<p align="right">(<a href="#readme-top">Volver al inicio</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[Go.dev]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[aws.com]: https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white
[Aws-url]: https://aws.amazon.com/es/
[dynamo.com]: https://img.shields.io/badge/Amazon%20DynamoDB-4053D6?style=for-the-badge&logo=Amazon%20DynamoDB&logoColor=white
[Dynamo-url]: https://aws.amazon.com/es/
[github.com]: https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white
[Github-url]: https://github.com/
[github-actions.com]: https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white
[GithubActions-url]: https://github.com/features/actions
