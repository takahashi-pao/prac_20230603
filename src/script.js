document.addEventListener('DOMContentLoaded', function() {
    document.getElementsByClassName('page-handler')[0].addEventListener('click', () => {
        fetch('./afternoon')
        .then(response => response.json())
        .then(data => {
            console.log(data)

            window.location.href = './afternoon';
        })
        .catch(error => {
            console.log(error);
        })
    });

    let elem = document.getElementsByClassName('page-handler-2')[0];
    if(elem != null){
        elem.addEventListener('click', () => {
            alert('clicked');
        });
    }

});
