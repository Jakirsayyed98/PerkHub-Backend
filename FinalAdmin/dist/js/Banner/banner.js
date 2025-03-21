const rowsPerPage = 25;
let currentPage = 1;
let Banners = []; // This holds all the Banners
let filteredBanners = []; // This is for the filtered Banners based on search

// Function to fetch Banners from the API
async function fetchBanners() {
    try {
        const categoryId = localStorage.getItem('category_id');
        console.log(categoryId)
        if (!categoryId) {
            throw new Error('Category ID is not available.');
        }
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/admin/get-banners/'+categoryId, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch Banners');
        }

        const data = await response.json();
        
        // Ensure that data.data is an array before using it
        Banners = Array.isArray(data.data) ? data.data : [];
        filteredBanners = [...Banners];  // Copy data to filteredBanners for later manipulation
        // alert('Banners fetched successfully');
        displayBanners();
    } catch (error) {
        console.error('Error fetching Banners:', error);
        alert('Error fetching Banners: ' + error.message);
    }
}

// Function to display Banners in the table
function displayBanners() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const BannersToDisplay = filteredBanners.slice(startIndex, endIndex);

    const tbody = document.getElementById('BannersTableBody');
    tbody.innerHTML = ''; // Clear the table body

    BannersToDisplay.forEach((Banners,index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
           <td>${index + 1}</td>
    
    <td><img src="${Banners.image}" alt="Game Image" width="75" height="75"/></td>
    <td>${Banners.name}</td>
    <td>${Banners.end_date}</td>
    <!-- Button for Status -->
    <td>
        <button class="btn btn-primary" id="Status" onclick="ActiveAndDeactive('${Banners.id}', ${Banners.status === true ? false : true}, 'Status')">
            ${Banners.status === true ?  'InActive' : 'Active'}
        </button>
    </td>

    <!-- Update Button -->
    <td>
                <button class="btn btn-danger" id="update-${Banners.id}">Update</button>
            </td>
    <!-- Delete Button -->
    <td>
        <button class="btn btn-danger" id="delete-${Banners.id}"">Delete</button>
    </td>
        `;
        tbody.appendChild(row);
        document.getElementById(`update-${Banners.id}`).addEventListener('click', function() {
            updateItem(Banners);
        });

        document.getElementById(`delete-${Banners.id}`).addEventListener('click', function() {
            DeleteBanners(Banners.id);
        });
    });

    createPagination();
}

// Store the item data in localStorage and navigate to page2.html for update
function updateItem(Banners) {
    // Directly store the Banners object in localStorage
    const itemData = { data: Banners };  // No need for decodeURIComponent
    localStorage.setItem('itemData', JSON.stringify(itemData));  // Store as stringified JSON
    window.location.href = './add_update_banner.html';  // Redirect to update page
}

// Function to create pagination buttons
function createPagination() {
    const totalPages = Math.ceil(filteredBanners.length / rowsPerPage);
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
            displayBannerss();
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
            displayBannerss();
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
            displayBannerss();
        }
    };
    nextButton.appendChild(nextLink);
    pagination.appendChild(nextButton);
}

// Function to filter users based on search input
function searchUsers() {
    const input = document.getElementById('searchInput').value.toLowerCase();
    filteredBanners = Banners.filter(Banners => {
        return (
            Banners.name.toLowerCase().includes(input) ||
            Banners.id.toString().includes(input) // You can add more fields for filtering if necessary
        );
    });
    currentPage = 1; // Reset to the first page after filtering
    displayBannerss();
}

// Event listener for search input
document.getElementById('searchInput').addEventListener('input', searchUsers);

// Initial fetch of Banners data when the script is loaded
fetchBanners();

// Function to fetch Banners from the API
async function DeleteBanners(bannerId) {
    try {
        console.log(bannerId)
       
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/admin/delete-banner/'+bannerId, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch Banners');
        }

        // Initial fetch of Banners data when the script is loaded
        alert('banner deleted successful!');
        fetchBanners();
    } catch (error) {
        console.error('Error fetching Banners:', error);
        alert('Error fetching Banners: ' + error.message);
    }
}

