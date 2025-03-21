const rowsPerPage = 25;
let currentPage = 1;
let affiliate_transactions = []; // This holds all the affiliate_transactions
let filteredaffiliate_transactions = []; // This is for the filtered affiliate_transactions based on search



// Function to fetch affiliate_transactions from the API
async function fetchaffiliate_transactions() {
    try {
        const token = localStorage.getItem('token'); // Get the token from localStorage
        
        const response = await fetch('http://localhost:4215/api/affiliate-transactions', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`, // Include the token in the request
                'Content-Type': 'application/json',
            },
            body: JSON.stringify( {"page_no":1,"limit":10}),
        });

        if (!response.ok) {
            throw new Error('Failed to fetch affiliate_transactions');
        }

        const data = await response.json();
        
        // Ensure that data.data is an array before using it
        affiliate_transactions = Array.isArray(data.data) ? data.data : [];
        filteredaffiliate_transactions = [...affiliate_transactions];  // Copy data to filteredaffiliate_transactions for later manipulation
        // alert('affiliate_transactions fetched successfully');
        displayaffiliate_transactions();
    } catch (error) {
        console.error('Error fetching affiliate_transactions:', error);
        alert('Error fetching affiliate_transactions: ' + error.message);
    }
}

// Function to display affiliate_transactions in the table
function displayaffiliate_transactions() {
    const startIndex = (currentPage - 1) * rowsPerPage;
    const endIndex = startIndex + rowsPerPage;
    const affiliate_transactionsToDisplay = filteredaffiliate_transactions.slice(startIndex, endIndex);

    const tbody = document.getElementById('affiliate_transactionTableBody');
    tbody.innerHTML = ''; // Clear the table body

    affiliate_transactionsToDisplay.forEach((affiliate_transaction,index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
    <td>${index + 1}</td>
    
    <td><img src="${affiliate_transaction}" alt="affiliate_transaction Image" width="75" height="75"/></td>
    <td>${affiliate_transaction.miniapp_name}</td>
    <td>${affiliate_transaction.sale_amount}</td>
    <td>${affiliate_transaction.order_id}</td>
    <td>${affiliate_transaction.commission}</td>
    <!-- Button for View -->
    <td>
        <button class="btn btn-primary" id="Status">
           View
        </button>
    </td>
`;

        tbody.appendChild(row);
    });

    createPagination();
}

// Store the item data in localStorage and navigate to page2.html for update
function updateItem(id, name, status) {
    const itemData = { id, name: decodeURIComponent(name), status: decodeURIComponent(status) };
    localStorage.setItem('itemData', JSON.stringify(itemData));
    window.location.href = 'AddAndEditMiniApp.html';
}

// Function to create pagination buttons
function createPagination() {
    const totalPages = Math.ceil(filteredaffiliate_transactions.length / rowsPerPage);
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
            displayaffiliate_transactions();
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
            displayaffiliate_transactions();
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
            displayaffiliate_transactions();
        }
    };
    nextButton.appendChild(nextLink);
    pagination.appendChild(nextButton);
}

// Function to filter users based on search input
function searchUsers() {
    const input = document.getElementById('searchInput').value.toLowerCase();
    filteredaffiliate_transactions = affiliate_transactions.filter(affiliate_transaction => {
        return (
            affiliate_transaction.name.toLowerCase().includes(input) ||
            affiliate_transaction.id.toString().includes(input) // You can add more fields for filtering if necessary
        );
    });
    currentPage = 1; // Reset to the first page after filtering
    displayaffiliate_transactions();
}


// Event listener for search input
document.getElementById('searchInput').addEventListener('input', searchUsers);

// Initial fetch of miniapp data when the script is loaded
fetchaffiliate_transactions();
