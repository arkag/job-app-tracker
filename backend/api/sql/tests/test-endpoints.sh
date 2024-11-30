#!/usr/bin/env bash

echo "___ Testing endpoints ___"
echo "___ GET jobapps ___"
curl -X GET http://localhost:8080/jobapps
echo ""
echo "___ POST jobapps/create ___"
curl -X POST -H "Content-Type: application/json" --data "$(cat create.json)" http://localhost:8080/jobapps/create
echo ""
echo "___ POST jobapps/update ___"
curl -X POST -H "Content-Type: application/json" --data "$(cat update.json)" http://localhost:8080/jobapps/update
echo ""
echo "___ POST jobapps/delete ___"
curl -X POST -H "Content-Type: application/json" --data "$(cat delete.json)" http://localhost:8080/jobapps/delete
echo ""
