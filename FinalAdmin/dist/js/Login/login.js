// Function to log in a user
async function loginUser(credentials) {
    try {
        const response = await fetch('http://localhost:4215/api/admin/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(credentials),
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
    const loginForm = document.getElementById('login-form');

    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault(); // Prevent default form submission

            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;

            try {
                // Call the login API
                const response = await loginUser({ email, password });

                if (response.token) {
                    // Save the token to localStorage
                    localStorage.setItem('token', response.token);
                    alert('Login successful!');
                    window.location.href = '../../index.html'; // Redirect to dashboard
                } else {
                    alert(response.message || 'Login failed. Please check your credentials.');
                }
            } catch (error) {
                console.error('Error during login:', error);
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