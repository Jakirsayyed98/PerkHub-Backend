const rowsPerPage = 25;
let currentPage = 1;
let banner_category = []; // This holds all the banner_category
let filteredbanner_category = []; // This is for the filtered banner_category based on search

// Function to fetch banner_category from the API
async function fetchBannerCategory() {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/admin/banner-category', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch banner category');
        }

        const data = await response.json();
        
        // Ensure that data.data is an array before using it
        banner_category = Array.isArray(data.data) ? data.data : [];
        filteredbanner_category = [...banner_category];  // Copy data to filteredbanner_category for later manipulation
        // alert('banner_category fetched successfully');
        displayBannerCategory();
    } catch (error) {
        console.error('Error fetching banner_category:', error);
        alert('Error fetching banner_category: ' + error.message);
    }
}

// Function to display banner_category in the table
function displayBannerCategory() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const banner_categoryToDisplay = filteredbanner_category.slice(startIndex, endIndex);

    const tbody = document.getElementById('BannerCategoryTableBody');
    tbody.innerHTML = ''; // Clear the table body

    banner_categoryToDisplay.forEach((banner_category,index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
           <td>${index + 1}</td>
    
    <td>${banner_category.title} ${banner_category.id}</td>
    
   
    <!-- Delete Button -->
    <td>
                <button class="btn btn-danger" id="update-${banner_category.id}">Update</button>
            </td>
    <!-- Delete Button -->
    <td>
        <button class="btn btn-danger" id="delete" onclick="window.location.href = ''">Delete</button>
    </td>
        `;
        tbody.appendChild(row);
        document.getElementById(`update-${banner_category.id}`).addEventListener('click', function() {
            updateItem(banner_category.id);
        });
    });

    createPagination();
}

// Store the item data in localStorage and navigate to page2.html for update
function updateItem(banner_categoryId) {
    // Directly store the banner_category object in localStorage
    localStorage.setItem('category_id', banner_categoryId);  // Use the appropriate variable for category_id

    window.location.href = './banner_list.html';  // Redirect to update page
}

// Function to create pagination buttons
function createPagination() {
    const totalPages = Math.ceil(filteredbanner_category.length / rowsPerPage);
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
            displayBannerCategory();
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
            displayBannerCategory();
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
            displayBannerCategory();
        }
    };
    nextButton.appendChild(nextLink);
    pagination.appendChild(nextButton);
}

// Function to filter users based on search input
function searchUsers() {
    const input = document.getElementById('searchInput').value.toLowerCase();
    filteredbanner_category = banner_category.filter(banner_category => {
        return (
            banner_category.name.toLowerCase().includes(input) ||
            banner_category.id.toString().includes(input) // You can add more fields for filtering if necessary
        );
    });
    currentPage = 1; // Reset to the first page after filtering
    displayBannerCategory();
}

// Event listener for search input
document.getElementById('searchInput').addEventListener('input', searchUsers);

// Initial fetch of banner_category data when the script is loaded
fetchBannerCategory();
async function AddBannerCategory(bannerCategory) {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const baseUrl = 'http://localhost:4215/api/admin/create-banner-category';
        const response = await fetch(baseUrl, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json', // Set content type to JSON
            },
            body: JSON.stringify(bannerCategory), // Convert data to JSON format before sending
        });

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error during API call:', error);
        throw error;
    }
}

// Handle form submission
document.addEventListener('DOMContentLoaded', () => {
    const miniappForm = document.getElementById('inputForm');

    if (miniappForm) {
        miniappForm.addEventListener('submit', async (e) => {
            e.preventDefault(); // Prevent default form submission

            const title = document.getElementById('inputText').value;

            try {
                // Package the form data into an object
                const bannerCategory = {
                    title: title,
                };

                // Call the AddBannerCategory API
                const response = await AddBannerCategory(bannerCategory);

                if (response.success) { // Assuming response.success indicates a successful request
                    alert('Banner added successfully');
                    fetchBannerCategory(); // Assuming this function reloads the banner categories

                    // Close the modal after successful submission
                    $('#inputModal').modal('hide'); // Hide the modal using Bootstrap's modal method
                } else {
                    alert(response.message || 'Failed to add banner category');
                }
            } catch (error) {
                console.error('Error during form submission:', error);
                alert('An error occurred. Please try again.');
            }
        });
    }
});
