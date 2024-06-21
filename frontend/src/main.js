import './style.css';
import { GetComputerInformation } from '../wailsjs/go/main/App'

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