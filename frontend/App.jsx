const coffeeItems = [
  { name: 'Espresso', description: 'Strong and smooth', price: '$3.50', tone: 'espresso', img: '/images/coffee/espresso.jpg' },
  { name: 'Latte', description: 'Milk and espresso', price: '$4.50', tone: 'latte', img: '/images/coffee/latte.jpg' },
  { name: 'Cappuccino', description: 'Foamy and rich', price: '$4.25', tone: 'cappuccino', img: '/images/coffee/capucino.jpg' }, // Проверь опечатку в названии файла (на скрине capucino.jpg с одной 'p')
  { name: 'Mocha', description: 'Chocolate comfort', price: '$4.75', tone: 'mocha', img: '/images/coffee/mocco.jpg' },       // На скрине файл называется mocco.jpg
  { name: 'Americano', description: 'Bold and clean', price: '$3.95', tone: 'americano', img: '/images/coffee/americano.jpg' },
  { name: 'Flat White', description: 'Velvety and balanced', price: '$4.60', tone: 'flatwhite', img: '/images/coffee/flatwhite.jpg' },
  { name: 'Moccachino', description: 'Creamy with a kick', price: '$4.10', tone: 'moccachino', img: '/images/coffee/moccachino.jpg' },
  { name: 'Cold Brew', description: 'Smooth and chilled', price: '$4.80', tone: 'coldbrew', img: '/images/coffee/coldbrew.jpg' },
];

function App() {
  return (
    <div className="page-shell">
      <header className="site-header">
        <div className="brand-block">
          <span className="brand-mark">Coffee.com</span>
          <p className="brand-tagline">Fresh roasts, creamy pours, and cozy blends.</p>
        </div>

        <nav className="top-nav" aria-label="Primary">
          <a href="#menu">Menu</a>
          <a href="#contact">Contact</a>
        </nav>
      </header>

      <main className="content-wrap">
        <section className="hero">
          <div className="hero-copy">
            <p className="eyebrow">Handcrafted coffee bar</p>
            <h1>Comfort in every cup.</h1>
            <p>
              A warm brown-and-white coffee shop landing page built in React,
              with a big logo, a clean menu grid, and space reserved for drink photos.
            </p>
          </div>

          <div className="hero-card">
            <span>Today’s roast</span>
            <strong>Dark chocolate, caramel, and toasted hazelnut.</strong>
          </div>
        </section>

        <section className="menu-section" id="menu">
          <div className="section-heading">
            <p className="eyebrow">Coffee menu</p>
            <h2>Signature drinks</h2>
          </div>

          <div className="coffee-grid">
            {coffeeItems.map((item) => (
              <article className="coffee-tile" key={item.name}>
                <div className={`photo-slot photo-${item.tone}`} aria-label={`${item.name} photo space`}>
                  <img src={item.img} alt={item.name} />
                </div>
                <footer className="tile-footer">
                  <div>
                    <h3>{item.name}</h3>
                    <p>{item.description}</p>
                  </div>
                  <strong>{item.price}</strong>
                </footer>
              </article>
            ))}
          </div>
        </section>
      </main>

      <footer className="site-footer" id="contact">
        <p>Open daily from 7:00 AM to 7:00 PM</p>
        <p>123 Roast Street, Coffee City</p>
      </footer>
    </div>
  );
}

export default App;