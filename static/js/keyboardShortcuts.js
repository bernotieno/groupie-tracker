// Function to navigate to a specific section based on keyboard input
function navigateToSection(sectionId) {
    const section = document.getElementById(sectionId);
    if (section) {
        section.scrollIntoView({ behavior: 'smooth' });
    }
}

// Function to navigate to the next or previous card
function navigateCards(direction) {
    const cards = document.querySelectorAll('.flip-card'); // Select all cards
    const currentIndex = Array.prototype.indexOf.call(cards, document.activeElement); // Get current focused card index
    
    let nextIndex;
    if (direction === 'next') {
        // Move to the next card, or wrap around to the first
        nextIndex = (currentIndex + 1) % cards.length;
    } else if (direction === 'prev') {
        // Move to the previous card, or wrap around to the last
        nextIndex = (currentIndex - 1 + cards.length) % cards.length;
    }

    // Focus the next or previous card
    cards[nextIndex].focus();
}


function showLoadingMessage() {
    const loadingMessage = document.getElementById('loadingMessage');

    // Create spinner element
    const spinner = document.createElement('div');
    spinner.classList.add('spinner');
    
    const cancelButton = document.createElement('button');
    cancelButton.innerText = 'Cancel';
    
    loadingMessage.innerHTML = 'Loading artist info...';
    
    // Insert the spinner before the text
    loadingMessage.insertBefore(spinner, loadingMessage.firstChild);
    
    // Add the Cancel button
    loadingMessage.appendChild(cancelButton);
    loadingMessage.style.display = "block";

    // Form reference and submit handler
    let form = null;
    let submitHandler = null;
    
    cancelButton.addEventListener('click', function() {
        hideLoadingMessage(); // Cancel the loading and hide the message
        // clearTimeout(loadingTimeout); 
        console.log("Passed")
        window.location.href = '/';
    });
}

// Hide loading message
function hideLoadingMessage() {
    const loadingMessage = document.getElementById('loadingMessage');
    if (loadingMessage) {
        loadingMessage.style.display = 'none';
    }
}

document.addEventListener('keydown', function(event) {

     if (event.code === 'ControlLeft') {
        navigateToSection('home');
    }

    if (event.code === 'ControlRight') {
        navigateToSection('featured-artists');
    }

    if (event.key === 'ArrowLeft') {
        navigateCards('prev');
    }

    if (event.key === 'ArrowRight') {
        navigateCards('next');
    }

    if (event.key === 'Enter' && document.activeElement.classList.contains('flip-card')) {
        event.preventDefault();
        showLoadingMessage();
        
        const focusedCard = document.activeElement;
        const form = focusedCard.closest('form');
        
        if (form) {
            const submitHandler = function() {
                hideLoadingMessage();
            };

            form.addEventListener('submit', submitHandler);

            setTimeout(() => {
                hideLoadingMessage();
                form.submit(); 
            }, 3000); 
        }
    }
});
