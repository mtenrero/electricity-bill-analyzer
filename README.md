# Electricity bill Analyzer

[![Go Report Card](https://goreportcard.com/badge/github.com/mtenrero/electricity-bill-analyzer)](https://goreportcard.com/report/github.com/mtenrero/electricity-bill-analyzer)
[![Build Status](https://travis-ci.org/mtenrero/electricity-bill-analyzer.svg?branch=master)](https://travis-ci.org/mtenrero/electricity-bill-analyzer)

This reposiory contains the code for the Golang Cloud Functions/Lambdas used to parse and analyze your real consumptions provided by your electricity distribution company and gives you practical information about hourly and daily usage.

This will help you in making clever decissions about your electricity provider promotions and product selection.

## Technical data

### parser

For now, it's only compatible with Spanish Iberdrola (i-DE) electricity provider, which allows the contract owner download a CSV file with the summary of the power consumed.

The package has a byte array as argument with the content of the previously mentioned CSV file

### cloud-functions

The backend is intended to be run in a serverless deployment, this package contains the cloud functions.