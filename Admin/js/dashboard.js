// Function to fetch total users from the API
async function dashboarddata() {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage

        const response = await fetch('http://localhost:4215/api/admin/dashboard-data', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch total users');
        }

        const data = await response.json();
        return data.data; // Assuming the API returns { totalUsers: 1234 }
    } catch (error) {
        console.error('Error fetching total users:', error);
        return null;
    }
}

// Function to update the total users in the DOM
async function updateTotalUsers() {
    const totalUsersElement = document.getElementById('total-users');
    const totalMiniapp = document.getElementById('minapp-count');
    const totalGame = document.getElementById('game-count');
    const adminName = document.getElementById('admin-name');

    if (totalUsersElement) {
        const totalCount = await dashboarddata();
        
        // if (totalUsers !== null) {
            totalUsersElement.textContent = totalCount.user_count; // Update the DOM
            totalMiniapp.textContent =  totalCount.minapp_count; // Update the DOM
            totalGame.textContent =  totalCount.games_count; // Update the DOM
            const name = localStorage.getItem('name');
            adminName.textContent =  name; // Update the DOM
        // } else {
        //     totalUsersElement.textContent = 'Failed to load'; // Show error message
        // }
    }
}

// Call the function to update total users when the page loads
document.addEventListener('DOMContentLoaded', () => {
    updateTotalUsers();
});