document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('toMain')?.addEventListener('click', setUserRating, false);
    document.getElementById('contentClick')?.addEventListener('click', contentOnClick, false);
    document.getElementById('getArticles')?.addEventListener('click', showListOnClick, false);

    chrome.storage.sync.get('userRating', function(result) {
        if (!result?.['userRating']) {
            const mainView = document.getElementById('mainView');
            const startingView = document.getElementById('startingView');
            mainView.style.display = 'none';
            startingView.style.display = 'block';
        }
    });

    function setUserRating() {
        // add a slider and allow user to set their value
        const startingSliderValue = document.getElementById("startingSlider")?.value;
        const startingToggleChecked = document.getElementById('startingToggle')?.checked;
        let userAlignment = startingToggleChecked ? 'na' : sliderValueToText(startingSliderValue); // save as the text representation of a political leaning
        chrome.storage.sync.set({'userRating': userAlignment}, function() {
            const mainView = document.getElementById('mainView');
            const startingView = document.getElementById('startingView');
            mainView.style.display = 'block';
            startingView.style.display = 'none';
        });
    }

    function contentOnClick() {
        chrome.tabs.query({currentWindow: true, active: true}, function (tabs) {
            chrome.tabs.sendMessage(tabs[0].id, {request: 'toContent'}, null, fromContent);
        });
    }

    function showListOnClick() {
        chrome.tabs.query({currentWindow: true, active: true}, function (tabs) {
            chrome.tabs.sendMessage(tabs[0].id,{request: 'getArticles'}, null, showArticleList);
        });
    }

    function fromContent(res) {
        if (res?.firstHeadingValue) {
            const headingDiv = document.createElement('div');
            headingDiv.textContent = res.firstHeadingValue;
            const br = document.createElement('br');
            document.body.appendChild(headingDiv);
            document.body.appendChild(br);
        }
        if (res?.url) {
            const urlDiv = document.createElement('div');
            urlDiv.textContent = res.url;
            const br = document.createElement('br');
            document.body.appendChild(urlDiv);
            document.body.appendChild(br);
        }
    }

    function showArticleList(res) {

        if (res?.listArticles) {
            const listArticles = res.listArticles;
            listArticles.forEach(a => {
                const anchorEl = document.createElement('a');
                anchorEl.textContent = a.title;
                anchorEl.href = a.url;
                anchorEl.onclick = () => openUrlInNewTab(a.url);
                document.body.appendChild(anchorEl);
            });
        }
    }

    function openUrlInNewTab(url) {
        chrome.tabs.create({ url: url });
    }

    var slider = document.getElementById("myRange");
    var output = document.getElementById("demo");
    output.innerHTML = sliderValueToText(slider.value); // Display the default slider value

    // Update the current slider value (each time you drag the slider handle)
    slider.oninput = function() {
        output.innerHTML = sliderValueToText(this.value);
    }

    // starting view slider:
    const startingSlider = document.getElementById("startingSlider");
    const startingSliderOutput = document.getElementById("startingSliderValue");
    startingSliderOutput.innerHTML = sliderValueToText(startingSlider.value); // Display the default slider value
    startingSlider.oninput = function() {
        startingSliderOutput.innerHTML = sliderValueToText(this.value);
    }

    // starting view toggle:
    const startingToggleValue = document.getElementById("startingToggle");
    startingToggleValue.addEventListener('change', function() {
        const slideContainer = document.getElementById('slideContainer');
        if(this.checked) {
            slideContainer.style.display = 'none';
        } else {
            slideContainer.style.display = 'block';
        }
    });
    startingSliderOutput.innerHTML = sliderValueToText(startingSlider.value); // Display the default slider value
    startingSlider.oninput = function() {
        startingSliderOutput.innerHTML = sliderValueToText(this.value);
    }

}, false);

function sliderValueToText(sliderValue) {
    let userAlignment = 'na';
    switch (sliderValue) {
        case '0':
            userAlignment = 'Left';
            break;
        case '1':
            userAlignment = 'Center-left';
            break;
        case '2':
            userAlignment = 'Center';
            break;
        case '3':
            userAlignment = 'Center-right';
            break;
        case '4':
            userAlignment = 'Right';
            break;
        default:
            break;
    }
    return userAlignment;
}
