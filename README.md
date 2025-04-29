## KYC Matching API
A Go-based API for matching KYC (Know Your Customer) data against a HarperDB database, built with the Echo framework and Dockerized for easy deployment.
## Prerequisites
Go: Version 1.23 or higher
Docker: For running the API and HarperDB
Git: For cloning the repository
Postman: For testing API endpoints
WSL Ubuntu 22.04.5 (or compatible Linux environment)
## Project Structure
├── README.md
├── .dockerignore
├── .gitignore
├── .github/workflows/docker-ci.yml
├── Dockerfile
├── config/config.go
├── database/database.go
├── go.mod
├── go.sum
├── handlers/kycHandler.go
├── main.go
├── models/models.go
├── repositories/harperdbRepository.go
├── repositories/kycRepository.go
├── server/server.go
├── usecases/kycUsecase.go
├── usecases/kycUsecaseImpl.go
## Setup Instructions
1. Clone the Repository
git clone https://github.com/balu6914/KYC-Match-API.git
cd KYC-Match-API

2. Set Up HarperDB
Run HarperDB in a Docker container:
docker run -d --name harperdb -p 9925:9925 -e HDB_ADMIN_USERNAME=HDB_ADMIN -e HDB_ADMIN_PASSWORD=password harperdb/harperdb

## Create a Docker network for communication between the API and HarperDB:
- docker network create kyc-network
- docker network connect kyc-network harperdb

- Verify HarperDB is running:
- docker ps

3. Configure Environment Variables
Create a .env file in the project root:
Add the following content:
HARPERDB_HOST=harperdb
HARPERDB_PORT=9925
HARPERDB_USERNAME=HDB_ADMIN
HARPERDB_PASSWORD=password
HARPERDB_SCHEMA=kyc_data

4. Build and Run the API
Build the Docker image:
- docker build -t balu1921/kyc-match-api:latest .

Run the API container on the kyc-network:
- docker run --rm -p 8080:8080 --env-file .env --network kyc-network --name kyc-api balu1921/kyc-match-api:latest

- Check logs to confirm the API is running:
- docker logs kyc-api

Expected Output:
Loaded Config: Host=harperdb, Port=9925, Username=HDB_ADMIN, Password=..., Schema=kyc_data
⇨ http server started on [::]:8080

5. Test the API with Postman
- Import the following Postman requests to test the /match endpoint:
## Non-Existent Phone Number
Method: POST
URL: http://localhost:8080/match
Headers:
Content-Type: application/json
x-correlator: b4333c46-49c0-4f62-80d7-f0ef930f1c46
Body:{
  "phoneNumber": "+99999999999",
  "idDocument": "12345678z"
}

Expected Response: 404 Not Found{
  "code": "IDENTIFIER_NOT_FOUND",
  "message": "No customer found for phoneNumber: +99999999999",
  "status": "404"
}

## Matching Request
Body:{
  "phoneNumber": "+34629255833",
  "idDocument": "66666666q",
  "givenName": "Federica",
  "familyName": "Sanchez Arjona",
  "name": "Federica Sanchez Arjona",
  "birthdate": "1978-08-22",
  "email": "abc@example.com"
}

Expected Response: 200 OK with "value": "true"

Non-Matching Request

Body:{
  "phoneNumber": "+34629255833",
  "idDocument": "12345678z",
  "givenName": "Maria",
  "familyName": "Gonzalez",
  "name": "Maria Gonzalez",
  "birthdate": "1980-01-01",
  "email": "maria@example.com"
}


Expected Response: 200 OK with "value": "false", "score": 85
## Invalid Request
Body:{
  "phoneNumber": ""
}


Expected Response: 400 Bad Request{
  "error": "at least one field besides phoneNumber must be provided"
}
## Docker Commands
- Build the image:docker build -t balu1921/kyc-match-api:latest .

- Run the API:docker run --rm -p 8080:8080 --env-file .env --network kyc-network --name kyc-api balu1921/kyc-match-api:latest

- Run HarperDB:docker run -d --name harperdb -p 9925:9925 -e HDB_ADMIN_USERNAME=HDB_ADMIN -e HDB_ADMIN_PASSWORD=password harperdb/harperdb

- Create Docker network:docker network create kyc-network
- docker network connect kyc-network harperdb

- Check logs:docker logs kyc-api
- Remove image:docker rmi balu1921/kyc-match-api:latest
## CI/CD with GitHub Actions
The .github/workflows/docker-ci.yml workflow:

Triggers: Runs on pushes to main, pull requests, or tags (v*).
Jobs:
test: Runs go test ./....
build-and-push: Builds and pushes the Docker image to balu1921/kyc-match-api for tags (e.g., v1.0.1).

To trigger a new image build:
git tag v1.0.1
git push origin v1.0.1

Check the workflow at: https://github.com/balu6914/KYC-Match-API/actions
Contributing

- Fork the repository.
- Create a feature branch (git checkout -b feature/your-feature).
- Commit changes (git commit -m "Add feature").
- Push to the branch (git push origin feature/your-feature).
- Open a pull request.# Test auto-build
