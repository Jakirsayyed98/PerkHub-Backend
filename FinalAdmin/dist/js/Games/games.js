const rowsPerPage = 25;
let currentPage = 1;
let games = []; // This holds all the games
let filteredUsers = []; // This is for the filtered games based on search

// Function to fetch games from the API
async function fetchGames() {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/getAllGames', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch games');
        }

        const data = await response.json();
        
        // Ensure that data.data is an array before using it
        games = Array.isArray(data.data) ? data.data : [];
        filteredUsers = [...games];  // Copy data to filteredUsers for later manipulation
        alert('games fetched successfully');
        displayGames();
    } catch (error) {
        console.error('Error fetching games:', error);
        alert('Error fetching games: ' + error.message);
    }
}

// Function to display games in the table
function displayGames() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const gamesToDisplay = filteredUsers.slice(startIndex, endIndex);

    const tbody = document.getElementById('GameTableBody');
    tbody.innerHTML = ''; // Clear the table body

    gamesToDisplay.forEach(game => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${game.index}</td>
            <td>${game.name}</td>
            <td>${game.status === '1' ? 'Active' : 'InActive'}</td>
            <td>
                <button id="addBrandBtn" class="btn btn-primary" onclick="window.location.href='AddAndEditMiniApp.html'">Update</button>
                <button id="addBrandBtn" class="btn btn-primary" onclick="window.location.href='AddAndEditMiniApp.html'">Delete</button>
            </td>
        `;
        tbody.appendChild(row);
    });

    createPagination();
}

// Store the item data in localStorage and navigate to page2.html for update
function updateItem(id, name, status) {
    const itemData = { id, name: decodeURIComponent(name), status: decodeURIComponent(status) };
    localStorage.setItem('itemData', JSON.stringify(itemData));
    window.location.href = 'AddAndEditMiniApp.html';
}

// Function to create pagination buttons
function createPagination() {
    const totalPages = Math.ceil(filteredUsers.length / rowsPerPage);
    const pagination = document.getElementById('pagination');
    pagination.innerHTML = ''; // Clear existing pagination

    // Prev button
    const prevButton = document.createElement('li');
    prevButton.classList.add('page-item');
    const prevLink = document.createElement('a');
    prevLink.classList.add('page-link');
    prevLink.href = '#';
    prevLink.textContent = '«';
    prevLink.onclick = () => {
        if (currentPage > 1) {
            currentPage--;
            displayGames();
        }
    };
    prevButton.appendChild(prevLink);
    pagination.appendChild(prevButton);

    // Page number buttons
    for (let i = 1; i <= totalPages; i++) {
        const pageButton = document.createElement('li');
        pageButton.classList.add('page-item');
        const pageLink = document.createElement('a');
        pageLink.classList.add('page-link');
        pageLink.href = '#';
        pageLink.textContent = i;
        pageLink.onclick = () => {
            currentPage = i;
            displayGames();
        };
        pageButton.appendChild(pageLink);
        pagination.appendChild(pageButton);
    }

    // Next button
    const nextButton = document.createElement('li');
    nextButton.classList.add('page-item');
    const nextLink = document.createElement('a');
    nextLink.classList.add('page-link');
    nextLink.href = '#';
    nextLink.textContent = '»';
    nextLink.onclick = () => {
        if (currentPage < totalPages) {
            currentPage++;
            displayGames();
        }
    };
    nextButton.appendChild(nextLink);
    pagination.appendChild(nextButton);
}

// Function to filter users based on search input
function searchGames() {
    const input = document.getElementById('searchInput').value.toLowerCase();
    filteredUsers = games.filter(user => {
        return (
            user.name.toLowerCase().includes(input) ||
            user.email.toLowerCase().includes(input) ||
            user.phone.includes(input)
        );
    });
    currentPage = 1; // Reset to the first page after filtering
    displayGames();
}

// Event listener for search input
document.getElementById('searchInput').addEventListener('input', searchUsers);

// Initial fetch of miniapp data when the script is loaded
fetchGames();
