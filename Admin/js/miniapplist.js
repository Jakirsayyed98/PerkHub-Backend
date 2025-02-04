const rowsPerPage = 10;
let currentPage = 1;
let miniapps = []; // This holds all the miniapps
let filteredUsers = []; // This is for the filtered miniapps based on search

// Function to fetch miniapps from the API
async function fetchMiniApps() {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/admin/AllMiniApps', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch miniapps');
        }

        const data = await response.json();
        
        // Ensure that data.data is an array before using it
        miniapps = Array.isArray(data.data) ? data.data : [];
        filteredUsers = [...miniapps];  // Copy data to filteredUsers for later manipulation
        alert('Miniapps fetched successfully');
        displayMiniApps();
    } catch (error) {
        console.error('Error fetching miniapps:', error);
        alert('Error fetching miniapps: ' + error.message);
    }
}

// Function to display miniapps in the table
function displayMiniApps() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const miniappsToDisplay = filteredUsers.slice(startIndex, endIndex);

    const tbody = document.getElementById('MiniAppTableBody');
    tbody.innerHTML = ''; // Clear the table body

    miniappsToDisplay.forEach(miniapp => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${miniapp.id}</td>
            <td>${miniapp.name}</td>
            <td>${miniapp.status === '1' ? 'Active' : 'InActive'}</td>
            <td>
            <button onclick="updateItem(${miniapp.id}, '${encodeURIComponent(miniapp.name)}', '${encodeURIComponent(miniapp.status)}')">Update</button>
            <button onclick="updateItem(${miniapp.id}, '${encodeURIComponent(miniapp.name)}', '${encodeURIComponent(miniapp.status)}')">Delete</button>
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
    const prevButton = document.createElement('a');
    prevButton.href = '#';
    prevButton.textContent = 'Prev';
    prevButton.classList.add('prev');
    prevButton.onclick = () => {
        if (currentPage > 1) {
            currentPage--;
            displayMiniApps();
        }
    };
    pagination.appendChild(prevButton);

    // Page number buttons
    for (let i = 1; i <= totalPages; i++) {
        const pageButton = document.createElement('a');
        pageButton.href = '#';
        pageButton.textContent = i;
        pageButton.onclick = () => {
            currentPage = i;
            displayMiniApps();
        };
        pagination.appendChild(pageButton);
    }

    // Next button
    const nextButton = document.createElement('a');
    nextButton.href = '#';
    nextButton.textContent = 'Next';
    nextButton.classList.add('next');
    nextButton.onclick = () => {
        if (currentPage < totalPages) {
            currentPage++;
            displayMiniApps();
        }
    };
    pagination.appendChild(nextButton);
}

// Function to filter users based on search input
function searchUsers() {
    const input = document.getElementById('searchInput').value.toLowerCase();
    filteredUsers = miniapps.filter(user => {
        return (
            user.name.toLowerCase().includes(input) ||
            user.email.toLowerCase().includes(input) ||
            user.phone.includes(input)
        );
    });
    currentPage = 1; // Reset to the first page after filtering
    displayMiniApps();
}

// Event listener for search input
document.getElementById('searchInput').addEventListener('input', searchUsers);

// Initial fetch of miniapp data when the script is loaded
fetchMiniApps();
