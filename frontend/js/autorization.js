function switchTab(tabName) {
    // Убираем активный класс со всех вкладок и форм
    document.querySelectorAll('.tab').forEach(tab => tab.classList.remove('active'));
    document.querySelectorAll('.form-section').forEach(section => section.classList.remove('active'));
    
    // Добавляем активный класс нужной вкладке и форме
    document.querySelector(`.tab:nth-child(${tabName === 'login' ? '1' : '2'}`).classList.add('active');
    document.getElementById(tabName).classList.add('active');
}

async function handleSubmit(event, type) {
    event.preventDefault();
    const form = event.target;
    const formData = {
        name: form.querySelector('[name="name"]').value,
        password: form.querySelector('[name="password"]').value
    };

    try {
        const endpoint = type === 'login' ? '/login' : '/registration';
        const response = await fetch(`http://localhost:8080/auth${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        });

        if (response.ok) {
            const result = await response.json();
            console.log('Success:', result);
            alert(type === 'login' ? 'Успешный вход!' : 'Успешная регистрация!');
        } else {
            const error = await response.json();
            console.error('Error:', error);
            alert(error.message || 'Произошла ошибка');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Ошибка соединения с сервером');
    }
}