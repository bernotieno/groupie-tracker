document.addEventListener('DOMContentLoaded', () => {
    // Select all forms in the featured artists section
    const forms = document.querySelectorAll('#featured-artists form');

    forms.forEach(form => {
        form.addEventListener('submit', function(event) {
            showLoadingMessage(); // Call the loading message function
            // Optionally, you can delay the form submission slightly to ensure the loading message is displayed.
            setTimeout(() => {
                // Allow the form to submit normally after a short delay
                this.submit();
            }, 100); // Adjust the delay if necessary
        });
    });
});

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
    
    // Insert the spinner before the text
    loadingMessage.insertBefore(spinner, loadingMessage.firstChild);
    loadingMessage.style.display = "block";
    
    // Add click event listener to the spinner to act as cancel
    spinner.addEventListener('click', function() {
        spinner.remove()
        hideLoadingMessage(); // Hide the loading message when spinner is clicked
        console.log("Cancel initiated by clicking the spinner");
        window.location.href = '/'; // Redirect to home (or adjust based on your need)
    });

    // Automatically hide the loading message after 5 seconds
    setTimeout( ()=>{
        spinner.remove()
        hideLoadingMessage
    }, 4000);
}

// Hide loading messagef4f4f4
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
            form.submit()
        }
    }
});
export {showLoadingMessage};