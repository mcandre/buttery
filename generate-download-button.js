const app_name = 'buttery';
const app_version = '0.0.34';
const u = 'https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.34';
const supportedPlatforms = [
    ['FreeBSD', 'arm64'],
    ['FreeBSD', 'x86_64'],
    ['Linux', 'arm64'],
    ['Linux', 'x86_64'],
    ['macOS', 'arm64'],
    ['macOS', 'x86_64'],
    ['Windows', 'arm64'],
    ['Windows', 'x86_64'],
];

function download_link(base_url, name, version, os, chipset) {
    switch (os) {
    case 'FreeBSD':
        let chipsetFreeBSD = '';

        switch (chipset) {
        case 'arm64':
            chipsetFreeBSD = 'aarch64';
            break;
        case 'x86_64':
            chipsetFreeBSD = 'amd64';
            break;
        }

        return `${base_url}/freebsd-${chipsetFreeBSD}/${name}-${version}.pkg`;
    case 'Linux':
        let chipsetAppImage = '';

        switch (chipset) {
        case 'arm64':
            chipsetAppImage = 'aarch64';
            break;
        case 'x86_64':
            chipsetAppImage = 'x86_64';
            break;
        }

        return `${base_url}/appimage/${name}-${version}-${chipsetAppImage}.AppImage`;
    case 'macOS':
        let chipsetMacOS = '';

        switch (chipset) {
        case 'arm64':
            chipsetMacOS = 'arm64';
            break;
        case 'x86_64':
            chipsetMacOS = 'x86_64';
            break;
        }

        return `${base_url}/macos/${name}-${chipsetMacOS}-${version}.pkg`;
    case 'Windows':
        let chipsetWindows = '';

        switch (chipset) {
        case 'arm64':
            chipsetWindows = 'Arm64';
            break;
        case 'x86_64':
            chipsetWindows = 'x64';
            break;
        }

        return `${base_url}/windows/${name}-${version}-${chipsetWindows}.msi`;
    }
}

new UAParser().getResult().withClientHints().then(result => {
    const os = (result.os?.name ?? 'unknown');
    const chipset = (result.cpu?.architecture ?? 'unknown');
    const platform = [os, chipset];
    const supportedDownloadP = document.getElementById('supported-download');
    const isSupported = supportedPlatforms.some(subArray =>
        subArray.length === platform.length &&
        subArray.every((value, index) => value === platform[index])
    );

    if (isSupported) {
        const a = document.createElement('a');
        a.id = 'download-button';
        a.href = download_link(u, app_name, app_version, os, chipset);
        a.classList.add(
            'btn',
            'btn-lg',
            'btn-warning',
            // 'btn-no-hover',
        );
        a.role = 'button';
        a.innerHTML = 'Download<i class="bi bi-download ms-2">';
        supportedDownloadP.appendChild(a);
        document.getElementById('platform').textContent = `For ${chipset} ${os}`;
    }
});
