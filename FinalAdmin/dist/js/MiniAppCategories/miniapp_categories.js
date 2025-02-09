const rowsPerPage = 25;
let currentPage = 1;
let miniappCategories = []; // This holds all the miniapp
let filteredMiniAppCategories = []; // This is for the filtered miniapp based on search

// Function to fetch miniapp from the API
async function fetchMiniAppCategories() {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/admin/get-category', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch miniapp');
        }

        const data = await response.json();
        
        // Ensure that data.data is an array before using it
        miniappCategories = Array.isArray(data.data) ? data.data : [];
        filteredMiniAppCategories = [...miniappCategories];  // Copy data to filteredMiniApp for later manipulation
        // alert('miniapp fetched successfully');
        displayMiniAppsCategories();
    } catch (error) {
        console.error('Error fetching miniapp Categories:', error);
        alert('Error fetching miniapp Categories: ' + error.message);
    }
}

// Function to display miniapp in the table
function displayMiniAppsCategories() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const miniappCategoriesToDisplay = filteredMiniAppCategories.slice(startIndex, endIndex);

    const tbody = document.getElementById('MiniAppTableBody');
    tbody.innerHTML = ''; // Clear the table body

    miniappCategoriesToDisplay.forEach((miniappCategories,index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
           <td>${index + 1}</td>
    
    <td><img src="${miniappCategories.image}" alt="Game Image" width="75" height="75"/></td>
    <td>${miniappCategories.name}</td>
   

    <!-- Button for Status -->
    <td>
        <button class="btn btn-primary" id="Status" onclick="ActiveAndDeactive('${miniappCategories.id}', ${miniappCategories.status === true ? false : true}, 'Status')">
            ${miniappCategories.status === true ?  'InActive' : 'Active'}
        </button>
    </td>

    <!-- Delete Button -->
    <td>
                <button class="btn btn-danger" id="update-${miniappCategories.id}">Update</button>
            </td>
    <!-- Delete Button -->
    <td>
        <button class="btn btn-danger" id="delete" onclick="window.location.href = './add_update_miniApp.html'">Delete</button>
    </td>
        `;
        tbody.appendChild(row);
        document.getElementById(`update-${miniappCategories.id}`).addEventListener('click', function() {
            updateItem(miniappCategories);
        });
    });

    createPagination();
}

// Store the item data in localStorage and navigate to page2.html for update
function updateItem(miniappCategories) {
    // Directly store the miniapp object in localStorage
    const itemData = { data: miniappCategories };  // No need for decodeURIComponent
    localStorage.setItem('itemData', JSON.stringify(itemData));  // Store as stringified JSON
    window.location.href = './add_update_miniApp_categories.html';  // Redirect to update page
}



// Function to create pagination buttons
function createPagination() {
    const totalPages = Math.ceil(filteredMiniAppCategories.length / rowsPerPage);
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
            displayMiniApps();
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
            displayMiniApps();
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
            displayMiniApps();
        }
    };
    nextButton.appendChild(nextLink);
    pagination.appendChild(nextButton);
}

// Function to filter users based on search input
function searchUsers() {
    const input = document.getElementById('searchInput').value.toLowerCase();
    filteredMiniAppCategories = miniapp.filter(miniapp => {
        return (
            miniapp.name.toLowerCase().includes(input) ||
            miniapp.id.toString().includes(input) // You can add more fields for filtering if necessary
        );
    });
    currentPage = 1; // Reset to the first page after filtering
    displayMiniApps();
}

// Event listener for search input
document.getElementById('searchInput').addEventListener('input', searchUsers);

// Initial fetch of miniapp data when the script is loaded
fetchMiniAppCategories();
