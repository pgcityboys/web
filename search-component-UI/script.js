// Mock danych tematów z wynikami
const mockTopics = [
    { id: 1, name: 'Matematyka', results: [
        { id: 101, name: 'Analiza matematyczna', description: 'Analiza matematyczna dla początkujących', date: '2024-06-20' },
        { id: 102, name: 'Algebra', description: 'Wprowadzenie do algebraicznych struktur', date: '2024-06-21' },
        { id: 103, name: 'Geometria', description: 'Podstawy geometrii analitycznej', date: '2024-06-22' }
    ] },
    { id: 2, name: 'Fizyka', results: [
        { id: 201, name: 'Mechanika klasyczna', description: 'Podstawy mechaniki Newtona', date: '2024-06-23' },
        { id: 202, name: 'Elektromagnetyzm', description: 'Teoria elektromagnetyczna Maxwella', date: '2024-06-24' }
    ] },
    { id: 3, name: 'Informatyka', results: [
        { id: 301, name: 'Programowanie w Pythonie', description: 'Kurs programowania w Pythonie', date: '2024-06-25' },
        { id: 302, name: 'Bazy danych', description: 'Wprowadzenie do baz danych', date: '2024-06-26' }
    ] },
    { id: 4, name: 'Biologia', results: [
        { id: 401, name: 'Biologia komórkowa', description: 'Podstawy biologii komórkowej', date: '2024-06-27' },
        { id: 402, name: 'Ekologia', description: 'Podstawy ekologii populacyjnej', date: '2024-06-28' }
    ] },
    { id: 5, name: 'Chemia', results: [
        { id: 501, name: 'Chemia organiczna', description: 'Wprowadzenie do chemii organicznej', date: '2024-06-29' },
        { id: 502, name: 'Chemia nieorganiczna', description: 'Podstawy chemii nieorganicznej', date: '2024-06-30' }
    ] },
    { id: 6, name: 'Historia', results: [
        { id: 601, name: 'Historia Polski', description: 'Podstawy historii Polski', date: '2024-07-01' },
        { id: 602, name: 'Historia Europy', description: 'Wprowadzenie do historii Europy', date: '2024-07-02' }
    ] },
    { id: 7, name: 'Literatura', results: [
        { id: 701, name: 'Literatura światowa', description: 'Wprowadzenie do literatury światowej', date: '2024-07-03' },
        { id: 702, name: 'Poezja', description: 'Analiza poezji romantycznej', date: '2024-07-04' }
    ] },
    { id: 8, name: 'Sztuka', results: [
        { id: 801, name: 'Historia sztuki', description: 'Podstawy historii sztuki', date: '2024-07-05' },
        { id: 802, name: 'Malarstwo', description: 'Techniki malarskie renesansu', date: '2024-07-06' }
    ] }
];

// Autocomplete dla pola wyszukiwania
const searchInput = document.getElementById('search-input');
const autocompleteDropdown = document.getElementById('autocomplete-dropdown');

searchInput.addEventListener('input', function() {
    const searchText = searchInput.value.toLowerCase();
    const matchedTopics = mockTopics.filter(topic =>
        topic.name.toLowerCase().includes(searchText)
    );
    displayAutocomplete(matchedTopics);
});

function displayAutocomplete(matchedTopics) {
    autocompleteDropdown.innerHTML = '';
    if (matchedTopics.length > 0) {
        autocompleteDropdown.style.display = 'block';
        matchedTopics.forEach(topic => {
            const option = document.createElement('div');
            option.classList.add('autocomplete-option');
            option.textContent = topic.name;
            option.addEventListener('click', function() {
                searchInput.value = topic.name;
                autocompleteDropdown.style.display = 'none';
                search(topic.results); // Przekazanie wyników dla wybranego tematu
            });
            autocompleteDropdown.appendChild(option);
        });
    } else {
        autocompleteDropdown.style.display = 'none';
    }
}

document.addEventListener('click', function(e) {
    if (!searchInput.contains(e.target) && !autocompleteDropdown.contains(e.target)) {
        autocompleteDropdown.style.display = 'none';
    }
});

// Funkcja wyszukiwania i wyświetlająca wyniki
function search(results) {
    const tbody = document.querySelector('#results-table tbody');
    tbody.innerHTML = '';
    results.forEach(result => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${result.name}</td>
            <td>${result.description}</td>
            <td>${result.date}</td>
        `;
        tbody.appendChild(row);
    });
}
