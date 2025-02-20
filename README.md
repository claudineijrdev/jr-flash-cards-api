# JR Flash Cards API

JR Flash Cards API is a backend service designed to support a flash cards application aimed at helping users learn English. This project is currently a work in progress and will serve both a web app and a mobile app (which are yet to be created). The API is developed in Go language and aims to grow into a complete service with more features over time.

## Features

- **Create Cards**: Users can create individual flash cards with questions and answers.
- **Create Decks**: Users can group flash cards into decks for organized learning sessions.
- **Play Decks**: Users can play through decks to test their knowledge and improve their language skills.

## Roadmap

Here is a tentative roadmap for the JR Flash Cards API:

- **Phase One**
  - Initial release with basic card and deck creation features
  - Basic user authentication and authorization

- **Phase Two**
  - Implement spaced repetition algorithm for better learning efficiency
  - Add support for multimedia flash cards (images, audio)

- **Phase Three**
  - Develop RESTful API endpoints for mobile and web app integration
  - Implement user progress tracking and analytics

- **Phase Four**
  - Introduce social features such as sharing decks and collaborative learning
  - Enhance security and performance optimizations

- **Phase Five**
  - Expand language support beyond English
  - Release mobile and web applications

## Getting Started

To get started with the JR Flash Cards API, follow these steps:

1. Clone the repository:
  ```sh
  git clone https://github.com/yourusername/jr-flash-cards-api.git
  ```
2. Navigate to the project directory:
  ```sh
  cd jr-flash-cards-api
  ```
3. Install dependencies:
  ```sh
  go mod tidy
  ```
4. Start the development server:
  ```sh
  go run main.go
  ```
