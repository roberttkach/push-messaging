# Distributed Messaging System

This project is a distributed messaging system deployed across multiple **Yandex Cloud** servers using Docker, Kubernetes, and Terraform. The system allows sending messages to recipients through various communication channels, including email, SMS, WhatsApp, Telegram, and Viber.

## Architecture

The system consists of the following components:

1. **Yandex Cloud Infrastructure**: Terraform is used to create and manage the infrastructure on Yandex Cloud, including 4 compute instances for hosting Docker containers.

2. **Docker**: The application is packaged into Docker containers for portability and scalability.

3. **Kubernetes**: Kubernetes is used for orchestrating and managing Docker containers on compute instances. Four containers are deployed on each server, totaling 16 containers.

4. **Go Application**: The main application, written in Go, is responsible for sending messages and interacting with various services (Email, SMS, Telegram, Whatsapp, Viber).

5. **PostgreSQL Database**: Used for storing recipient information.

## Deployment

1. **Yandex Cloud Infrastructure**: Run `terraform apply` to create compute instances and necessary resources on Yandex Cloud.

2. **Docker and Kubernetes**:
   - Build the Docker image using `docker build`.
   - Deploy a Kubernetes cluster on the compute instances by running `kubeadm init` on one of the servers.
   - Join the remaining servers to the cluster using `kubeadm join`.
   - Apply the Kubernetes manifests (`deployment.yaml`) to deploy the 16 application containers.

3. **Application Configuration**: Set the required environment variables for database access, messaging service credentials, and other configurations.

4. **Access the Application**: Obtain the public IP address of one of the compute instances and open the application's web interface to send messages.

## Testing

## Continuous Integration and Delivery (CI/CD)

This project incorporates a CI/CD pipeline for automated testing, building, and deployment. The `tests.go` file contains Go tests that are executed as part of the CI/CD process.

## Contributing

If you find a bug or would like to propose an improvement, please create a pull request. We welcome contributions!

## License

This project is distributed under the MIT License.
