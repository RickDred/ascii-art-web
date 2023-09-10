# ascii-art-web

Ascii-art-web is a project that involves creating and running a server to provide a web GUI (graphical user interface) version of the previous project, ascii-art. This web application allows users to generate ASCII art banners using different styles such as shadow, standard, and thinkertoy.

# Description

The Ascii-art-web project aims to provide a user-friendly interface for generating ASCII art banners. Users can select from various banner styles and input their desired text, after which the server processes the data and returns the corresponding ASCII art. The project is built using Go for the server-side logic and HTML templates for the frontend.

# Usage: How to Run

Follow these steps to run the Ascii-art-web project:

    Clone the Repository:
git clone [repository_url]
cd ascii-art-web

    Run the following command to start the server:

go run main.go

The server will start on http://localhost:8000.

    Access the Web GUI:

Open your web browser and go to http://localhost:8000. You will be presented with the main page of the Ascii-art-web application.

# Implementation Details: Algorithm

The Ascii-art-web project is implemented using Go for the server logic and HTML templates for the frontend. Here's a high-level overview of the algorithm:

    Server Initialization:

    The Go server is initialized using the Gin framework, which provides a lightweight and efficient HTTP server.

    Routing:

    The server defines two main routes:

        GET /: This route renders the main page of the web GUI. It uses Go templates to display the HTML content, including the input form for text and banner selection.

        POST /ascii-art: This route handles the form submission from the main page. It receives user input (text and banner style), processes it to generate the corresponding ASCII art, and then displays the result either on a new page (/ascii-art) or appends it to the home page.

    ASCII Art Generation:

    The server processes the user's input and generates the ASCII art based on the selected banner style (shadow, standard, thinkertoy). The specific algorithms for ASCII art generation are implemented based on the chosen styles.

    Displaying Results:

    The generated ASCII art is then displayed either on a new page (/ascii-art) or appended to the home page, depending on the chosen approach. This is achieved using Go templates to dynamically render the HTML content with the generated ASCII art.

    Frontend Interaction:

    The main page includes a form with dropdown options for selecting the banner style and an input field for entering the desired text. Users submit the form to the server, triggering the generation of ASCII art.

By following these implementation details, the Ascii-art-web project provides a web-based graphical interface for creating ASCII art banners with different styles.