// Function to log in a user
async function AddMiniAppCategories(miniAppData) {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const baseUrl = 'http://localhost:4215/api/admin/create-category';
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
            formData.append('description', formData.get('description'));
            formData.append('image', document.getElementById('image').files[0]);
            formData.append('status', formData.get('status') ? 'true' : 'false');
            formData.append('homepage_visible', formData.get('homepage_visible') ? 'true' : 'false');

            console.log(formData); // This will log the FormData object for debugging

            try {
                // Call the AddMiniAppCategories API with FormData
                const response = await AddMiniAppCategories(formData);

                if (response.ok) {
                    alert('Saved successful!');
                    window.location.href = './miniapp_categories.html'; // Redirect to dashboard
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
