#!/bin/bash

set -a
source .env
set +a

function saveFAQsToJson() {
    local response="$1"

    # Define the JSON file path
    local jsonFile="./json/faq-response.json"

    echo "$response" | jq . >$jsonFile

    if [ $? -ne 0 ]; then
        echo "Error: Failed to save FAQs to JSON."
        return 1
    fi

    echo "FAQs saved to $jsonFile"
}

function fetchFAQS() {
    local response=$(curl -X GET $BACKEND_URL:$BACKEND_PORT/faqs)
    echo "$response"
}

function addFAQ() {
    read -p "Enter the question: " question
    read -p "Enter the answer: " answer

    local response=$(curl -X POST $BACKEND_URL:$BACKEND_PORT/faqs \
        -H "Content-Type: $CONTENT_TYPE" \
        -d "{
    \"Question\": \"$question\",
    \"Answer\": \"$answer\"
}")

    echo $response

}

function deleteAllFAQs() {
    local response=$(curl -X DELETE $BACKEND_URL:$BACKEND_PORT/faqs)
    if [[ $? -ne 0 ]]; then
        echo "Error: Failed to fetch FAQs."
        return 1
    fi
    rm -rf ./json/faq-response.json
}

function deployToGoogleCloudRun() {
    local id="knowledge-manager-68472"
    docker build -t gcr.io/$id/faq-service .
    # gcloud auth login
    # gcloud config set project $id
    docker push gcr.io/$id/faq-service
    gcloud run deploy faq-service \
        --image gcr.io/$id/faq-service \
        --platform managed \
        --region us-central1 \
        --allow-unauthenticated

}

function runDockerContainerLatest() {
    docker run -e BACKEND_PORT=$BACKEND_PORT -p $BACKEND_PORT:$BACKEND_PORT $DOCKER_IMAGE_NAME:latest
}

function startFrontend() {
    npm --prefix ./frontend run dev
}

function startBackend() {
    go run ./main.go
}

function detectOSPlatform() {
    # Detect the platform
    platform=$(uname -m)

    case $platform in
    x86_64)
        os_platform="linux/amd64"
        ;;
    arm64)
        os_platform="linux/amd64"
        ;;
    aarch64)
        os_platform="linux/arm64"
        ;;
    armv7l)
        os_platform="linux/arm/v7"
        ;;
    *)
        echo "Unsupported architecture: $platform"
        exit 1
        ;;
    esac
}

function mainMenu() {
    # Start the CLI menu
    echo "**********************************"
    echo "         SELECT AN OPTION         "
    echo "**********************************"
    echo "1. Add new FAQ."
    echo "2. Save FAQs to JSON."
    echo "3. Delete all FAQs from database."
    echo "4. Start frontend."
    echo "5. Start backend."
    echo "6. Deploy to Google Cloud Run."
    echo "7. Run docker container (latest)."
}

function main() {
    detectOSPlatform
    mkdir -p ./json
    mainMenu

    read -p "Please enter a number [1-7]: " option

    case $option in
    1)
        addFAQ
        response=$(fetchFAQS)
        saveFAQsToJson "$response"
        exit 0
        ;;
    2)
        response=$(fetchFAQS)
        saveFAQsToJson "$response"
        exit 0
        ;;
    3)
        deleteAllFAQs
        response=$(fetchFAQS)
        saveFAQsToJson "$response"
        exit 0
        ;;
    4)
        startFrontend
        exit 0
        ;;
    5)
        startBackend
        exit 0
        ;;
    6)
        deployToGoogleCloudRun
        exit 0
        ;;
    7)
        runDockerContainerLatest
        exit 0
        ;;
    *)
        echo "Exiting"
        exit 0
        ;;
    esac
}

main
