// Function to log in a user
async function AddBanner(BannerData) {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const baseUrl = 'http://localhost:4215/api/admin/create-banner';
        const response = await fetch(baseUrl, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
            },
            body: BannerData, // Pass FormData directly
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
    const BannerForm = document.getElementById('banner-form');

    if (BannerForm) {
        BannerForm.addEventListener('submit', async (e) => {
            e.preventDefault(); // Prevent default form submission
            const categoryId = localStorage.getItem('category_id');
            // Create a new FormData object from the form
            const formData = new FormData(); // Initialize FormData object

            // Collecting the text data from form fields
            formData.append('id', document.getElementById('id').value);
            formData.append('name', document.getElementById('name').value);
            formData.append('banner_category_id', categoryId);
            formData.append('url', document.getElementById('url').value);
            formData.append('start_date', document.getElementById('start_date').value);
            formData.append('end_date', document.getElementById('end_date').value);
          
            // Collecting the status as a boolean (active/inactive)
            formData.append('status', document.getElementById('status').checked ? 'true' : 'false');
          
            // Collecting the image file (if selected)
            const imageFile = document.getElementById('image').files[0];
            if (imageFile) {
              formData.append('image', imageFile);
            }
          
            try {
                // Call the AddBanner API with FormData
                const response = await AddBanner(formData);

                if (response.ok) {
                    alert('Saved successful!');
                    window.location.href = '../banner_list.html'; // Redirect to dashboard
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
