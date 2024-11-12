// Select main elements from the provided HTML
const searchInput = document.getElementById('searchInput');
const artistForms = document.querySelectorAll('.artist-grid form');
const searchResult = document.getElementById('searchResult');

// Create range filter sliders for Creation Date and First Album Date
const creationDateRange = createRangeFilter('Creation Date', 1960, 2023);
const firstAlbumDateRange = createRangeFilter('First Album Date', 1960, 2023);

// Create checkbox filter for Number of Members
const memberCountFilters = createMemberCheckboxes([1, 2, 3, 4, 5, 6, 7, 8]);

// Append filters to the search form container in the provided HTML
document.querySelector('.hero-content').append(creationDateRange, firstAlbumDateRange, memberCountFilters);

// Create location search input with label
const locationContainer = document.createElement('div');
locationContainer.classList.add('location-filter');

const locationLabel = document.createElement('label');
locationLabel.innerText = 'Location:';

const locationInput = document.createElement('input');
locationInput.type = 'text';
locationInput.id = 'locationInput';
locationInput.placeholder = 'Search by Location';

// Append label and input to location container, then to the search form container
locationContainer.append(locationLabel, locationInput);
document.querySelector('.hero-content').appendChild(locationContainer);

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

// Search form submission handling to show results based on search input
function handleSearch(event) {
  event.preventDefault();
  const searchTerm = searchInput.value.trim();

  if (searchTerm) {
    searchResult.style.display = 'block';
    filterArtists();
  }
}
