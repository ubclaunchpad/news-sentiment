chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    if (message.request === 'toContent') {
        const heading = document.getElementsByTagName("h1");
        const firstHeadingValue = heading[0]?.innerText;
        const url = window.location.href;
        sendResponse({firstHeadingValue: firstHeadingValue, url: url});
    } else if (message.request === 'getDummy') {
        // do a fetch call based on the current article here
        // maybe something like fetch('http:///articles/{id}').then(response => sendResponse({listArticles: the-response.json()-array}));
        // where id is the id or url for the current article

        // send dummy list for now:
        const listArticles = [
            {
                ID: 0,
                TITLE: 'We Tried It: What I Learned About My Rescue Puppy Through a Dog DNA Screening',
                URL: 'https://people.com/pets/we-tried-it-dog-dna-test-for-rescue-puppy/'
                // Votes []Vote ignore votes for now
            },
            {
                ID: 1,
                TITLE: 'We Tried It: What I Learned About My Rescue Puppy Through a Dog DNA Screening',
                URL: 'https://people.com/pets/we-tried-it-dog-dna-test-for-rescue-puppy/'
                // Votes []Vote ignore votes for now
            },
            {
                ID: 2,
                TITLE: 'We Tried It: What I Learned About My Rescue Puppy Through a Dog DNA Screening',
                URL: 'https://people.com/pets/we-tried-it-dog-dna-test-for-rescue-puppy/'
                // Votes []Vote ignore votes for now
            },
            {
                ID: 3,
                TITLE: 'We Tried It: What I Learned About My Rescue Puppy Through a Dog DNA Screening',
                URL: 'https://people.com/pets/we-tried-it-dog-dna-test-for-rescue-puppy/'
                // Votes []Vote ignore votes for now
            },
            {
                ID: 4,
                TITLE: 'We Tried It: What I Learned About My Rescue Puppy Through a Dog DNA Screening',
                URL: 'https://people.com/pets/we-tried-it-dog-dna-test-for-rescue-puppy/'
                // Votes []Vote ignore votes for now
            },
            {
                ID: 5,
                TITLE: 'We Tried It: What I Learned About My Rescue Puppy Through a Dog DNA Screening',
                URL: 'https://people.com/pets/we-tried-it-dog-dna-test-for-rescue-puppy/'
                // Votes []Vote ignore votes for now
            },
            {
                ID: 6,
                TITLE: 'We Tried It: What I Learned About My Rescue Puppy Through a Dog DNA Screening',
                URL: 'https://people.com/pets/we-tried-it-dog-dna-test-for-rescue-puppy/'
                // Votes []Vote ignore votes for now
            }];
        sendResponse({listArticles: listArticles});
    }
    return true;
});
