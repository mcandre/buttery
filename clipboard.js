document.querySelectorAll('pre').forEach((pre) => {
    // 1. Create a relative wrapper around the <pre>
    const wrapper = document.createElement('div');
    wrapper.className = 'pre-wrapper';
    pre.parentNode.insertBefore(wrapper, pre);
    wrapper.appendChild(pre);

    // 2. Create the copy button
    const btn = document.createElement('button');
    btn.className = 'btn btn-dark copy-btn bi bi-clipboard';

    btn.addEventListener('click', async () => {
        const text = pre.querySelector('code')?.innerText || pre.innerText;
        await navigator.clipboard.writeText(text);

        btn.classList.remove('bi-clipboard');
        btn.textContent = 'Copied';

        setTimeout(() => {
            btn.textContent = '';
            btn.classList.add('bi-clipboard');
        }, 2000);
    });

    // 3. Append the button to the wrapper instead of the <pre>
    wrapper.appendChild(btn);
});
