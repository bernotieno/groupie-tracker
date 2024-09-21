let debounceTimer;
const searchForm = document.getElementById('searchForm');
const searchInput = document.getElementById('searchInput');
const searchResult = document.getElementById('searchResult');
const debounce = (func, wait) => {
    let timeout;

    return function (...args) {
        const later = () => {
            clearTimeout(timeout);
            func.apply(this, args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait); 
    };
};
const debouncedSearch = debounce(performSearch,500)
const handleClick = () => {
    console.log('Search !!');
};
const de = debounce(handleClick, 500)
searchInput.addEventListener('input', () => {
    debouncedSearch()
    de();
});


async function performSearch() {
    const query = searchInput.value.trim();
    
    if (query.length > 0) {
        try {
            const results = await fetchFromServer('/search?q=' + encodeURIComponent(query));
            const uniqueResults = new Set();

            // Filter the results to keep only unique ones
            const filteredResults = results.filter(result => {
                const isDuplicate = uniqueResults.has(result.Result);
                uniqueResults.add(result.Result);
                return !isDuplicate;
            });

            // Check if results are empty, null, or not an array
            if (!results || !Array.isArray(results) || filteredResults.length === 0) {
                displayErrorMessage(); // Ensure the error message is displayed
            } else {
                displayResults(filteredResults);
            }
        } catch (error) {
            displayErrorMessage(); // Handle errors during the fetch
        }
    } else {
        clearResults(); // Clear results if the input is empty
    }
}

function displayErrorMessage() {
    clearResults();
    const errorDiv = document.createElement('div');
    errorDiv.className = 'error-message';
    errorDiv.textContent = `Result Not Found`;
    searchResult.appendChild(errorDiv);
    searchResult.style.display = 'block';
}

function displayResults(results) {
    clearResults();
    
        const resultsList = document.createElement('ul');
        resultsList.className = 'results-list';
        results.forEach(result => {
            
            // Create a clickable list item
            const li = document.createElement('li');
            li.className = 'result-item';

            const artistLink = document.createElement('a');
            artistLink.textContent = result.Result;
            artistLink.href = '#';
            
            // Add a click event listener to the button
            artistLink.addEventListener('click', (e) => {
                e.preventDefault();
                submitFormWithArtist(result.ArtistName);
            });
            
            li.appendChild(artistLink);
            resultsList.appendChild(li);
        });
        searchResult.appendChild(resultsList);
    
    searchResult.style.display = 'block';
}

function submitFormWithArtist(artistName) {
    // Create a form element
    const form = document.createElement('form');
    form.method = 'POST'; // or 'GET' depending on your needs
    form.action = '/artistInfo'; // the URL where you want to submit the form

    // Create a hidden input field to hold the artist name
    const input = document.createElement('input');
    input.type = 'hidden';
    input.name = 'ArtistName';
    input.value = artistName;
    
    // Append the input field to the form
    form.appendChild(input);
    
    // Append the form to the body and submit it
    document.body.appendChild(form);
    form.submit();
}

function clearResults() {
    searchResult.innerHTML = '';
    searchResult.style.display = 'none';
}

// Initial search on page load if there's a query in the URL
window.addEventListener('load', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const initialQuery = urlParams.get('q');
    if (initialQuery) {
        searchInput.value = initialQuery;
        debouncedSearch();performSearch
    }
});