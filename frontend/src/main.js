import './style.css';
import { GetComputerInformation, GetConfig, SaveConfig, CheckProcessingStatus, Authenticate } from '../wailsjs/go/main/App'

document.addEventListener('DOMContentLoaded', async () => {
    const info = await GetComputerInformation();

    createTableRow('.computer-info', 'User', info.user);
    createTableRow('.computer-info', 'Hostname', info.hostname);
    createTableRow('.computer-info', 'MAC Address', info.macAddress);

    document.body.classList.remove('hidden');
})

function createTableRow(table, a, b) {
    const row = document.createElement('tr');
    const c1 = document.createElement('td');
    const c2 = document.createElement('td');
    const c3 = document.createElement('td');

    c1.innerText = a;
    c2.innerText = ':';
    c3.innerText = b;

    row.append(c1, c2, c3);

    document.querySelector(table).append(row);
}

document.querySelector('#save-btn').addEventListener('click', async () => {
    const tenantUrl = document.querySelector('#tenant-url').value;
    const userEmail = document.querySelector('#user-email').value;

    const config = await GetConfig();
    if (tenantUrl != "") config.tenantUrl = tenantUrl;
    if (userEmail != "") config.userEmail = userEmail;
    await SaveConfig(config);

    CheckProcessingStatus().then(setStatus);
    Authenticate().then(code => {
        alert('Auth finished with status code: ' + code);
    })

    alert("Please check your phone.");
})

function setStatus(status) {
    document.querySelector('#status').innerText = status
    
    const iconEl = document.querySelector('#status-icon');
    if (status == 'authenticated') {
        iconEl.classList.remove('bi-unlock');
        iconEl.classList.add('bi-lock');
    } else {
        iconEl.classList.remove('bi-lock');
        iconEl.classList.add('bi-unlock');
    }
}