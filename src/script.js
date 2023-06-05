document.addEventListener('DOMContentLoaded', function() {
    let elem1 = document.getElementsByClassName('page-handler')[0];
    if(elem1 != null){
        elem1.addEventListener('click', () => {
            fetch('./afternoon')
            .then(response => response.json())
            .then(data => {
                
    
                window.location.href = './afternoonPage';
                console.log(data)
            })
            .catch(error => {
                console.log(error);
            })
        });
    }


    let elem2 = document.getElementsByClassName('page-handler-2')[0];
    if(elem2 != null){
        elem2.addEventListener('click', () => {
            alert('clicked');
        });
    }

});
