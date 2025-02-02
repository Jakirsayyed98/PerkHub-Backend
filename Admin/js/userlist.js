const rowsPerPage = 10;
let currentPage = 1;
let users = [];
let filteredUsers = [];

// Function to fetch users from the API
async function fetchUsers() {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/admin/user-list', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch users');
        }

        const data = await response.json();
        users = data.data; // Assuming the API returns users in data.data
        filteredUsers = [...users]; // Initialize filteredUsers with all users
        displayUsers(); // Display the users on the page
    } catch (error) {
        console.error('Error fetching users:', error);
        alert('Error fetching users');
    }
}

// Function to display users in the table
function displayUsers() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const usersToDisplay = filteredUsers.slice(startIndex, endIndex);

    const tbody = document.getElementById('userTableBody');
    tbody.innerHTML = ''; // Clear the table body

    usersToDisplay.forEach(user => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${user.name}</td>
            <td>${user.name}</td>
            <td>${user.email}</td>
            <td>${user.number}</td>
            <td>${user.gender}</td>
            <td>${user.verified == '1' ? 'Verified' : 'Unverified'}</td>
        `;
        tbody.appendChild(row);
    });

    createPagination();
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
            displayUsers();
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
            displayUsers();
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
            displayUsers();
        }
    };
    pagination.appendChild(nextButton);
}

// Function to filter users based on search input
function searchUsers() {
    const input = document.getElementById('searchInput').value.toLowerCase();
    filteredUsers = users.filter(user => {
        return (
            user.name.toLowerCase().includes(input) ||
            user.email.toLowerCase().includes(input) ||
            user.phone.includes(input)
        );
    });
    currentPage = 1; // Reset to the first page
    displayUsers();
}

// Event listener for search input
document.getElementById('searchInput').addEventListener('input', searchUsers);

// Initial fetch of user data when the script is loaded
fetchUsers();
