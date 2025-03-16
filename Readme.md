# NEPSE Navigator

<p align="center"><img src="https://socialify.git.ci/rohankarn35/NEPSE-NAVIGATOR/image?font=Jost&language=1&name=1&owner=1&pattern=Overlapping+Hexagons&theme=Light"></p>

## Description

NEPSE Navigator is a comprehensive tool designed to provide detailed insights and data on the Nepal Stock Exchange (NEPSE). This project is built using Go and leverages GraphQL to serve market data stored in a MongoDB database. The data is periodically collected and updated by a separate server using cron jobs.

## Features

- **Market Movers**: Fetch data on the top gainers and losers in the market.
- **NEPSE Index**: Retrieve the current NEPSE index values.
- **IPO/FPO Alerts**: Get information on upcoming and ongoing IPOs/FPOs.
- **Market Status**: Check the current status of the market.

## Prerequisites

- Go (version 1.16 or higher)
- MongoDB (version 4.4 or higher)
- Git

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/rohankarn35/NEPSE-NAVIGATOR
   cd NEPSE-NAVIGATOR
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Configure MongoDB:
   - Ensure MongoDB is running and accessible.
   - Set up the MongoDB connection string in the configuration file or environment variables.

## Configuration

Set up the following environment variables:

- `MONGODB_URI`: MongoDB connection string (e.g., `mongodb://localhost:27017/nepse`)

## Usage

1. Run the server:
   ```sh
   go run server.go
   ```
2. Sample GraphQL queries:
   - Fetch Market Movers:
     ```graphql
     {
       marketMovers {
         symbol
         price
         change
       }
     }
     ```
   - Fetch NEPSE Index:
     ```graphql
     {
       nepseIndex {
         index
         change
       }
     }
     ```

## Schema Overview

The GraphQL schema exposes the following types of data:

- **MarketMovers**: Information on top gainers and losers.
- **NEPSEIndex**: Current NEPSE index values.
- **IPOAlerts**: Details on IPOs and FPOs.
- **MarketStatus**: Information about the market status.

## Architecture

The system consists of two main components:

1. **Cron-job Server**: Periodically fetches NEPSE market data and stores it in MongoDB.
2. **GraphQL Server**: Queries the MongoDB database to serve market data via a GraphQL API.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## Community

Join our Telegram channel for updates and discussions: [NEPSE Navigator Telegram Channel](https://t.me/nepsenavigator)

## License

This project is licensed under the MIT License. See the [LICENSE](./License.md) file for details.
