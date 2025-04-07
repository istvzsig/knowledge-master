#!/bin/bash

BASE_URL="http://localhost:8080"
CONTENT_TYPE="application/json"

# Create the json directory if it doesn't exist
mkdir -p ./json

function saveFAQsToJson() {
    local response="$1"

    # Define the JSON file path
    local jsonFile="./json/faq-response.json"

    # Check if the response is empty
    if [[ -z "$response" ]]; then
        # Create an empty JSON array if the file does not exist
        if [[ ! -f "$jsonFile" ]]; then
            echo "[]" >"$jsonFile"
            echo "Created an empty JSON array in $jsonFile"
        else
            echo "Error: No response provided."
            return 1
        fi
    else
        # Check if the JSON file exists
        if [[ -f "$jsonFile" ]]; then
            # Read the existing content
            existingContent=$(<"$jsonFile")

            # Combine existing content with new response
            combinedContent=$(echo "$existingContent" | jq --argjson newData "$response" '. += [$newData]')

            # Save the combined content back to the JSON file
            echo "$combinedContent" >"$jsonFile"
        else
            # If the file does not exist, create it with the new response in an array
            echo "[$response]" | jq . >"$jsonFile"
        fi
    fi

    if [ $? -ne 0 ]; then
        echo "Error: Failed to save FAQs to JSON."
        return 1
    fi

    echo "FAQs saved to $jsonFile"
}

function fetchFAQS() {
    local response=$(curl -X GET http://localhost:8080/faqs)
    if [[ "$response" == "null" ]]; then
        echo "Response is null"
        rm -rf ./json/faq-response.json
        return 1
    fi
    saveFAQsToJson "$response"
}

function addFAQ() {
    read -p "Enter the question: " question
    read -p "Enter the answer: " answer

    local response=$(curl -X POST $BASE_URL/faqs \
        -H "Content-Type: $CONTENT_TYPE" \
        -d "{
    \"Question\": \"$question\",
    \"Answer\": \"$answer\"
}")
    saveFAQsToJson "$response"
}

function deleteAllFAQs() {
    local response=$(curl -X DELETE $BASE_URL/faqs)
    if [[ $? -ne 0 ]]; then
        echo "Error: Failed to fetch FAQs."
        return 1
    fi
    rm -rf ./json/faq-response.json
}

echo "**********************************"
echo "         SELECT AN OPTION         "
echo "**********************************"
echo "1. Add new FAQ."
echo "2. Get FAQS and save to JSON."
echo "3. Delete all FAQs."
read -p "Please enter a number: " option

case $option in
1)
    addFAQ
    exit 0
    ;;
2)
    fetchFAQS
    exit 0
    ;;
3)
    deleteAllFAQs
    exit 0
    ;;
*)
    echo "Exiting"
    exit 0
    ;;
esac
