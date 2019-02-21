package core

//go:generate python ./script/golang-json-converter/model_generator.py --input ./json/ --output ./models/ --package models --caution
//go:generate goimports -w ./models/
