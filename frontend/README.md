# Run on Chrome
Go to
```
chrome://extensions/
```

in your chrome browser.

Turn on developer mode through the Developer mode toggle on the upper right-hand side.

Click "Load unpacked" in the upper left-hand side.

Choose the news-sentiment root folder.

Click on the refresh icon if you make a change to the file.

You should be able to use the extension through the browser now.

------------------
As of Nov 22:
Currently, the values are stored as either 'na', 'left', 'center-left', 'center',
'center-right', or 'right' in chrome.storage.sync (the developer panel doesn't 
seem show this - you might have to download the Storage Area Explorer extension
(https://chrome.google.com/webstore/detail/storage-area-explorer/ocfjjjjhkpapocigimmppepjgfdecjkb?hl=en)) 
to see or delete this data (there is no way currently to change a user's preference
 from the GUI).
