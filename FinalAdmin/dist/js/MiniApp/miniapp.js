const rowsPerPage = 25;
let currentPage = 1;
let miniapp = []; // This holds all the miniapp
let filteredMiniApp = []; // This is for the filtered miniapp based on search

// Function to fetch miniapp from the API
async function fetchMiniApp() {
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
            throw new Error('Failed to fetch miniapp');
        }

        const data = await response.json();
        
        // Ensure that data.data is an array before using it
        miniapp = Array.isArray(data.data) ? data.data : [];
        filteredMiniApp = [...miniapp];  // Copy data to filteredMiniApp for later manipulation
        // alert('miniapp fetched successfully');
        displayMiniApps();
    } catch (error) {
        console.error('Error fetching miniapp:', error);
        alert('Error fetching miniapp: ' + error.message);
    }
}

// Function to display miniapp in the table
function displayMiniApps() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const miniappToDisplay = filteredMiniApp.slice(startIndex, endIndex);

    const tbody = document.getElementById('MiniAppTableBody');
    tbody.innerHTML = ''; // Clear the table body

    miniappToDisplay.forEach((miniapp,index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
           <td>${index + 1}</td>
    
    <td><img src="${miniapp.icon}" alt="Game Image" width="75" height="75"/></td>
    <td>${miniapp.name}</td>
    
    <!-- Button for Popular -->
    <td>
        <button class="btn btn-primary" id="Popular" onclick="ActiveAndDeactive('${miniapp.id}', ${miniapp.popular === true ? false : true}, 'Popular')">
            ${miniapp.popular === true ? 'InActive' : 'Active'}
        </button>
    </td>

    <!-- Button for Trending -->
    <td>
        <button class="btn btn-primary" id="Trending" onclick="ActiveAndDeactive('${miniapp.id}', ${miniapp.trending === true ? false : true}, 'Trending')">
            ${miniapp.trending === true ?  'InActive' : 'Active'}
        </button>
    </td>

    <!-- Button for Status -->
    <td>
        <button class="btn btn-primary" id="Status" onclick="ActiveAndDeactive('${miniapp.id}', ${miniapp.status === true ? false : true}, 'Status')">
            ${miniapp.status === true ?  'InActive' : 'Active'}
        </button>
    </td>

    <!-- Delete Button -->
    <td>
                <button class="btn btn-danger" id="update-${miniapp.id}">Update</button>
            </td>
    <!-- Delete Button -->
    <td>
        <button class="btn btn-danger" id="delete-${miniapp.id}">Delete</button>
    </td>
        `;
        tbody.appendChild(row);
        document.getElementById(`update-${miniapp.id}`).addEventListener('click', function() {
            updateItem(miniapp);
        });
        document.getElementById(`delete-${miniapp.id}`).addEventListener('click', function() {
            DeleteMiniApp(miniapp.id);
        });
    });

    createPagination();
}

// Store the item data in localStorage and navigate to page2.html for update
function updateItem(miniapp) {
    // Directly store the miniapp object in localStorage
    const itemData = { data: miniapp };  // No need for decodeURIComponent
    localStorage.setItem('itemData', JSON.stringify(itemData));  // Store as stringified JSON
    window.location.href = './add_update_miniApp.html';  // Redirect to update page
}



// Function to create pagination buttons
function createPagination() {
    const totalPages = Math.ceil(filteredMiniApp.length / rowsPerPage);
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
    filteredMiniApp = miniapp.filter(miniapp => {
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
fetchMiniApp();


// Function to fetch miniapp from the API
async function DeleteMiniApp(miniAppId) {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage
        alert('Are yousure you want to delete?');
        const response = await fetch('http://localhost:4215/api/admin/delete-miniapp/'+miniAppId, {
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
        miniapp = Array.isArray(data.data) ? data.data : [];
        filteredMiniApp = [...miniapp];  // Copy data to filteredMiniApp for later manipulation
        // alert('miniapp fetched successfully');
        fetchMiniApp();
    } catch (error) {
        console.error('Error fetching miniapp:', error);
        alert('Error fetching miniapp: ' + error.message);
    }
}