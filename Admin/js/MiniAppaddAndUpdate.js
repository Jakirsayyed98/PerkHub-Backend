// Function to handle form submission with multipart/form-data (file upload)
async function submitMiniAppForm(formData) {
    try {
        const token = localStorage.getItem('token'); 
        const response = await fetch('http://localhost:4215/api/admin/create-miniapp', {
            method: 'POST',
            headers: {
                // Set only Authorization header (do not manually set Content-Type for FormData)
                'Authorization': `Bearer ${token}`, // Include the token in the request
            },
            body: formData, // Send the formData (which includes files)
        });

        const data = await response.json(); // Parse the JSON response
        return data;
    } catch (error) {
        console.error('Error during API call:', error);
        throw error;
    }
}

// Handle form submission
document.addEventListener('DOMContentLoaded', () => {
    const miniAppForm = document.getElementById('miniAppForm');

    if (miniAppForm) {
        miniAppForm.addEventListener('submit', async (e) => {
            e.preventDefault(); // Prevent default form submission

            // Create FormData object from the form
            const formData = new FormData(miniAppForm);

            // Optional: add additional fields or data if needed
            // formData.append('extraField', 'value');

            try {
                // Call the submitMiniAppForm API function
                const response = await submitMiniAppForm(formData);

                if (response.success) {
                    alert('MiniApp created successfully!');
                    window.location.href = '/miniapp.html'; // Redirect to another page
                } else {
                    alert(response.message || 'Failed to create MiniApp');
                }
            } catch (error) {
                console.error('Error during form submission:', error);
                alert('An error occurred. Please try again.');
            }
        });
    }
});


function validateToken() {
    const token = localStorage.getItem('token');

    if (!token) {
        return false;
    }

    // Decode the token to check expiration (if applicable)
    const tokenData = JSON.parse(atob(token.split('.')[1])); // Decode the token payload
    const now = Math.floor(Date.now() / 1000); // Current time in seconds

    if (tokenData.exp && tokenData.exp < now) {
        // Token is expired
        localStorage.removeItem('token'); // Clear the expired token
        return false;
    }

    return true;
}