#!/bin/bash

git config --local user.name "anxielray"
git config --local user.email "anxielworld@gmail.com"
git config --local credential.helper store
echo "Updating documentation..."
echo "starting the Git staging process..."
# cd Lexer
git add add.sh
git commit -m "[Script]: Add a script for the staging process"
git push
