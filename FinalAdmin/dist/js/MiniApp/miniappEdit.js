const rowsPerPage = 25;
let currentPage = 1;
let miniappCategories = []; // This holds all the miniapp
let filteredMiniAppCategories = []; // This is for the filtered miniapp based on search

// Function to log in a user
async function AddMiniApp(miniAppData) {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const baseUrl = 'http://localhost:4215/api/admin/create-miniapp';
        const response = await fetch(baseUrl, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
            },
            body: miniAppData, // Pass FormData directly
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
    const miniappForm = document.getElementById('miniapp-form');

    if (miniappForm) {
        miniappForm.addEventListener('submit', async (e) => {
            e.preventDefault(); // Prevent default form submission

            // Create a new FormData object from the form
            const formData = new FormData(miniappForm);

            // Collect all form data values into FormData (for file uploads too)
            formData.append('id', formData.get('id'));
            formData.append('name', formData.get('name'));
            formData.append('miniapp_category_id', formData.get('miniapp_category_id'));
            formData.append('miniapp_subcategory_id', formData.get('miniapp_subcategory_id'));
            formData.append('description', formData.get('description'));
            formData.append('howitswork', formData.get('howitswork'));
            formData.append('about', formData.get('about'));
            formData.append('cashback_terms', formData.get('cashback_terms'));
            formData.append('cashback_rates', formData.get('cashback_rates'));
            formData.append('url_type', formData.get('url_type'));
            formData.append('cb_percentage', formData.get('cb_percentage'));
            formData.append('url', formData.get('url'));
            formData.append('label', formData.get('label'));
            formData.append('macro_publisher', formData.get('macro_publisher'));
            formData.append('status', formData.get('status') ? 'true' : 'false');
            formData.append('cb_active', formData.get('cb_active') ? 'true' : 'false');
            formData.append('popular', formData.get('popular') ? 'true' : 'false');
            formData.append('trending', formData.get('trending') ? 'true' : 'false');
            formData.append('top_cashback', formData.get('top_cashback') ? 'true' : 'false');

            // Files are automatically handled by FormData, no need to manually append
            formData.append('banner', document.getElementById('banner').files[0]);
            formData.append('logo', document.getElementById('logo').files[0]);
            formData.append('icon', document.getElementById('icon').files[0]);

            console.log(formData); // This will log the FormData object for debugging

            try {
                // Call the AddMiniApp API with FormData
                const response = await AddMiniApp(formData);

                if (response.ok) {
                    alert('Saved successful!');
                    window.location.href = './miniapp.html'; // Redirect to dashboard
                } else {
                    alert(response.message || 'Save failed. Please check your credentials.');
                }
            } catch (error) {
                console.error('Error during form submission:', error);
                alert('An error occurred. Please try again.');
            }
        });
    }
});

fetchMiniAppCategories();
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
function displayMiniAppsCategories(type) {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const miniappCategoriesToDisplay = filteredMiniAppCategories.slice(startIndex, endIndex);

    const select = document.getElementById('miniapp_category_id');
    
    select.innerHTML = ''; // Clear the select options
   
    miniappCategoriesToDisplay.forEach((miniappCategory) => {
        
        const option = document.createElement('option');
        option.value = miniappCategory.id;
        option.textContent = miniappCategory.name;
        select.appendChild(option)
    });
}