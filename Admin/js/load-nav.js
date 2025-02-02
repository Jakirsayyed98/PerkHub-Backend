// load-nav.js
document.addEventListener("DOMContentLoaded", function () {
    // Fetch the navigation bar
    fetch("navbar.html")
        .then((response) => response.text())
        .then((data) => {
            // Insert the navigation bar into the placeholder
            document.getElementById("navbar-placeholder").innerHTML = data;
        })
        .catch((error) => console.error("Error loading the navigation bar:", error));
});