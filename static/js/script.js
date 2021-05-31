document.addEventListener('DOMContentLoaded', () => {
    const menuItems = ['fire_page_admin', 'fire_phrase_admin'];
    menuItems.forEach((i) => {
        let button = document.getElementById(i)
        button.addEventListener('click', () => {
            let modal = document.getElementById(button.dataset.target)
            let modal_back = document.querySelector(`#${button.dataset.target} .modal-background`)
            modal_back.addEventListener('click', ()=>{
                modal.classList.toggle('is-active')
            })
            modal.classList.toggle('is-active')
        })
    });
})