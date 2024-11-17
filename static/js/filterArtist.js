// Select main elements from the provided HTML
const artistForms = document.querySelectorAll('.artist-grid form');


// Create range filter sliders for Creation Date and First Album Date
// Create filter container div that floats on top right
// Create filter toggle button
const filterToggleBtn = document.createElement('button');
filterToggleBtn.innerHTML = '🔍 Filters';
filterToggleBtn.style.cssText = `
    position: fixed;
    top: 15px;
    right: 12px;
    padding: 6px 14px;
    background: transparent;
    border: 1px solid rgba(255, 215, 0, 0.9);
    border-radius: 5px;
    cursor: pointer;
    font-weight: bold;
    z-index: 1001;
    transition: all 0.3s ease;
    font-size: 14px;
    color: white;
`;

filterToggleBtn.addEventListener('mouseover', () => {
    filterToggleBtn.style.background = 'rgba(255, 215, 0, 1)';
    filterToggleBtn.style.transform = 'scale(1.05)';
    filterToggleBtn.style.color = 'black';
});

filterToggleBtn.addEventListener('mouseout', () => {
    filterToggleBtn.style.background = 'transparent';
    filterToggleBtn.style.transform = 'scale(1)';
    filterToggleBtn.style.color = 'white';
});

// Add media query for mobile
const mediaQuery = window.matchMedia('(max-width: 768px)');
const handleMobileStyles = (e) => {
    if (e.matches) {
        filterToggleBtn.style.padding = '6px 12px';
        filterToggleBtn.style.fontSize = '12px';
        filterToggleBtn.style.height = '25px';
        filterToggleBtn.style.top = '20px';
    }
};
mediaQuery.addListener(handleMobileStyles);
handleMobileStyles(mediaQuery);

const filterContainer = document.createElement('div');
filterContainer.style.cssText = `
    position: fixed;
    top: 10px;
    right: -100%;
    padding: 15px;
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    border: 2px solid #ffd700;
    border-radius: 10px;
    z-index: 1000;
    width: 90%;
    max-width: 400px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    transition: right 0.3s ease-in-out;
    margin-top: 40px;
    @media (max-width: 768px) {
        width: 85%;
        padding: 10px;
        margin-top: 35px;
    }
    @media (max-width: 480px) {
        width: 80%;
        padding: 8px;
    }
`;

const creationDateRange = createRangeFilter('Creation Date', 1960, 2023);
const firstAlbumDateRange = createRangeFilter('First Album Date', 1960, 2023);

// Create checkbox filter for Number of Members
const memberCountFilters = createMemberCheckboxes([1, 2, 3, 4, 5, 6, 7, 8]);

// Create location search input with label
const locationContainer = document.createElement('div');
locationContainer.classList.add('location-filter');

const locationLabel = document.createElement('label');
locationLabel.innerText = 'Location:';
locationLabel.style.cssText = `
    display: block;
    margin-bottom: 5px;
    color: #333;
    font-weight: bold;
    font-size: 14px;
    @media (max-width: 768px) {
        font-size: 12px;
    }
`;

const locationInput = document.createElement('input');
locationInput.type = 'text';
locationInput.id = 'locationInput';
locationInput.placeholder = 'Filter by Location';
locationInput.style.cssText = `
    width: 100%;
    padding: 8px 12px;
    border: 1px solid #ddd;
    border-radius: 5px;
    background: rgba(255, 255, 255, 0.9);
    font-size: 14px;
    transition: border-color 0.3s ease;
    @media (max-width: 768px) {
        padding: 6px 10px;
        font-size: 12px;
    }
    &:focus {
        outline: none;
        border-color: #ffd700;
        box-shadow: 0 0 5px rgba(255, 215, 0, 0.3);
    }
`;

// Add label and input to location container
locationContainer.appendChild(locationLabel);
locationContainer.appendChild(locationInput);

// Style the filter elements
[creationDateRange, firstAlbumDateRange, memberCountFilters, locationContainer].forEach(element => {
    element.style.cssText += `
        margin-bottom: 15px;
        padding: 10px;
        background: rgba(255, 255, 255, 0.1);
        border-radius: 8px;
        @media (max-width: 768px) {
            margin-bottom: 10px;
            padding: 8px;
        }
    `;
});

// Style checkboxes and range inputs
memberCountFilters.querySelectorAll('input[type="checkbox"]').forEach(checkbox => {
    checkbox.style.cssText = `
        margin-right: 8px;
        cursor: pointer;
        accent-color: #ffd700;
        width: 16px;
        height: 16px;
        border: 2px solid #ddd;
        border-radius: 3px;
        transition: all 0.2s ease;
        position: relative;
        vertical-align: middle;
        
        &:hover {
            border-color: #ffd700;
            box-shadow: 0 0 3px rgba(255, 215, 0, 0.3);
        }
        
        &:checked {
            background-color: #ffd700;
            border-color: #ffd700;
        }
        
        @media (max-width: 768px) {
            width: 14px;
            height: 14px;
            margin-right: 6px;
        }
    `;
});

memberCountFilters.querySelectorAll('label').forEach(label => {
    label.style.cssText = `
        margin-right: 12px;
        color: #333;
        cursor: pointer;
        font-size: 14px;
        @media (max-width: 768px) {
            font-size: 12px;
            margin-right: 8px;
        }
    `;
});

// Append all filters to the floating container
filterContainer.append(creationDateRange, firstAlbumDateRange, memberCountFilters, locationContainer);

// Add the floating container and toggle button to the body
document.body.appendChild(filterToggleBtn);
document.body.appendChild(filterContainer);

// Toggle filter panel visibility
let isFilterVisible = false;
filterToggleBtn.addEventListener('click', () => {
    isFilterVisible = !isFilterVisible;
    filterContainer.style.right = isFilterVisible ? '10px' : '-100%';
    filterToggleBtn.innerHTML = isFilterVisible ? '✕ Close' : '🔍 Filters';
});

// Add event listeners to the range sliders, checkboxes, and search inputs
creationDateRange.addEventListener('input', filterArtists);
firstAlbumDateRange.addEventListener('input', filterArtists);
searchInput.addEventListener('input', filterArtists);
locationInput.addEventListener('input', filterArtists);
memberCountFilters.addEventListener('change', filterArtists);

// Filter function based on all criteria
function filterArtists() {
  const searchTerm = searchInput.value.toLowerCase();
  const locationTerm = locationInput.value.toLowerCase();
  const creationDateMin = parseInt(creationDateRange.querySelector('.min').value);
  const creationDateMax = parseInt(creationDateRange.querySelector('.max').value);
  const firstAlbumMin = parseInt(firstAlbumDateRange.querySelector('.min').value);
  const firstAlbumMax = parseInt(firstAlbumDateRange.querySelector('.max').value);

  const selectedMembers = Array.from(memberCountFilters.querySelectorAll('input[type="checkbox"]:checked'))
                               .map(checkbox => parseInt(checkbox.value));

  // Loop through each artist form and apply filters
  artistForms.forEach(form => {
    const artistName = form.getAttribute('data-artist-name').toLowerCase();
    const creationDate = parseInt(form.getAttribute('data-creation-date'));
    const firstAlbumDate = parseInt(form.getAttribute('data-first-album-date'));
    const numberOfMembers = parseInt(form.getAttribute('data-number-of-members'));

    // Get artist concert locations and split by commas
    const concertLocations = form.getAttribute('data-concert-locations')
                                  .toLowerCase()
                                  .split(', ')
                                  .map(location => location.trim());

    // Check if any location contains the location search term
    const matchesLocation = locationTerm === '' || 
                            concertLocations.some(location => location.includes(locationTerm));

    // Check other filters
    const matchesSearch = artistName.includes(searchTerm);
    const matchesCreationDate = creationDate >= creationDateMin && creationDate <= creationDateMax;
    const matchesFirstAlbumDate = firstAlbumDate >= firstAlbumMin && firstAlbumDate <= firstAlbumMax;
    const matchesMembers = selectedMembers.length === 0 || selectedMembers.includes(numberOfMembers);

    // Display or hide the form based on all filter conditions
    if (matchesSearch && matchesLocation && matchesCreationDate && matchesFirstAlbumDate && matchesMembers) {
      form.style.display = 'block';
    } else {
      form.style.display = 'none';
    }
  });
}

// Helper function to create a range filter (date range slider)
function createRangeFilter(label, min, max) {
  const container = document.createElement('div');
  container.classList.add('range-filter');

  const labelElement = document.createElement('label');
  labelElement.innerText = `${label}:`;

  const minInput = document.createElement('input');
  minInput.type = 'number';
  minInput.className = 'min';
  minInput.value = min;
  minInput.min = min;
  minInput.max = max;

  const maxInput = document.createElement('input');
  maxInput.type = 'number';
  maxInput.className = 'max';
  maxInput.value = max;
  maxInput.min = min;
  maxInput.max = max;

  container.append(labelElement, minInput, document.createTextNode(' - '), maxInput);
  return container;
}

// Helper function to create checkboxes for number of members
function createMemberCheckboxes(memberOptions) {
  const container = document.createElement('div');
  container.classList.add('checkbox-filter');

  const labelElement = document.createElement('label');
  labelElement.innerText = 'Number of Members:';
  container.append(labelElement);

  memberOptions.forEach(count => {
    const checkbox = document.createElement('input');
    checkbox.type = 'checkbox';
    checkbox.value = count;
    checkbox.id = `members-${count}`;

    const label = document.createElement('label');
    label.innerText = count;
    label.htmlFor = `members-${count}`;

    container.append(checkbox, label);
  });

  return container;
}

