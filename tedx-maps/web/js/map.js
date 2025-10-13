const map = L.map('map').setView([43.238949, 76.889709], 13);

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
  maxZoom: 19,
  attribution: '&copy; OpenStreetMap contributors'
}).addTo(map);

async function loadPoints() {
  try {
    const res = await fetch('/api/points');
    const data = await res.json();
    data.forEach(p => {
      L.marker([p.latitude, p.longitude])
        .addTo(map)
        .bindPopup(`<b>${p.name}</b><br>${p.description}`);
    });
  } catch (e) {
    console.error('Ошибка загрузки точек:', e);
  }
}

loadPoints();
