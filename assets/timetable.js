var meetingsData = {};

var calendar = new Calendar("calendarContainer", "medium",
                    [ "Monday", 3 ],
                    [ "#9828b5", "#802398", "#ffffff", "#521A57" ],
                    {
                        days: [ "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday",  "Saturday" ],
                        months: [ "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December" ],
                        indicator: true,
                        placeholder: "<button onclick=openPopup()>Custom Placeholder</button>"
                    });

let fetchMeetings = () => {
    const apiUrl = 'http://localhost:2137/api/meetings/matematyka';
    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            var meetings = data;
            meetingsData = {};

            meetings.forEach(function(meeting) {
                if (!meetingsData[meeting.year]) {
                    meetingsData[meeting.year] = {};
                }
                if (!meetingsData[meeting.year][meeting.month]) {
                    meetingsData[meeting.year][meeting.month] = {};
                }
                if (!meetingsData[meeting.year][meeting.month][meeting.day]) {
                    meetingsData[meeting.year][meeting.month][meeting.day] = [];
                }
                
                meetingsData[meeting.year][meeting.month][meeting.day].push({
                    startTime: meeting.startTime,
                    endTime: meeting.endTime,
                    text: meeting.category,
                    link: meeting.link
                });
            });

            console.log("meetingsData po zgrupowaniu:", meetingsData);
            var organizer = new Organizer("organizerContainer", calendar, meetingsData);
        })
        .catch(error => {
            console.error('Wystąpił problem z pobraniem danych:', error);
        });
}

fetchMeetings();

function openPopup() {
    document.getElementById("popup-form").style.display = "block";
}

function closePopup() {
    document.getElementById("popup-form").style.display = "none";
    document.getElementById("eventForm").reset();
}

    document.getElementById("eventForm").addEventListener("submit", function(event) {
        event.preventDefault();
        var startTime = document.getElementById("startTime").value;
        var endTime = document.getElementById("endTime").value;
        var category = document.getElementById("category").value;
        var link = document.getElementById("link").value;
        var day = document.getElementById("day").value;
        var month = document.getElementById("month").value;
        var year = document.getElementById("year").value;

        var meetingData = {
            startTime: startTime,
            endTime: endTime,
            category: category,
            link: link,
            day: day,
            month: month,
            year: year
        };
        sendData(meetingData);
        closePopup();
    });

function sendData(data) {
    var apiUrl = 'http://localhost:2137/api/meetings/create/meeting';
    var fetchOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    };

    fetch(apiUrl, fetchOptions)
        .then(function(response) {
        fetchMeetings();
        return response.json();
    })
}