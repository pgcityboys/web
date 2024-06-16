var meetingsData = {};

var day = 15;
var month = 6;
var year = 2024;

var calendar = new Calendar("calendarContainer", "medium",
                    [ "Monday", 3 ],
                    [ "#d579ec", "#ca5ee6", "#ffffff", "#521A57" ],
                    {
                        days: [ "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday",  "Saturday" ],
                        months: [ "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December" ],
                        indicator: true,
                        placeholder: "<button class=addMeetingBtn onclick=openPopup()>Add meeting</button>"
                    });


let fetchMeetings = (category) => {
    const apiUrl = 'http://localhost:2137/api/meetings/' + category;
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
                    link: "/meeting/?id=" + meeting.id
                });
            });

            console.log("meetingsData po zgrupowaniu:", meetingsData);

            const elementsToDelete = document.querySelectorAll(".cjslib-events");
            elementsToDelete.forEach(element => {
                element.remove();
            });
            

            var organizer = new Organizer("organizerContainer", calendar, meetingsData);

            const radios = document.querySelectorAll('input[name="calendarContainer-day-radios"]');

            radios.forEach(radio => {
                radio.addEventListener('change', handleRadioChange);
            });

            function handleRadioChange(event) {
                console.log('Zmieniono wybór na: ' + event.target.value);
                console.log('ID klikniętego przycisku radiowego: ' + (Number.parseInt(event.target.id.split('-').slice(-1)) - 5));
                day = Number.parseInt(event.target.id.split('-').slice(-1)) - 5;
            }
        })
        .catch(error => {
            console.error('Wystąpił problem z pobraniem danych:', error);
        });
}

fetchMeetings("any");

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
        var link = "link";

        var meetingData = {
            startTime: startTime,
            endTime: endTime,
            category: category,
            link: link,
            day: day.toString(),
            month: month.toString(),
            year: year.toString()
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
        fetchMeetings("any");
        return response.json();
    })
}