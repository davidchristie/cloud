import ProductSearch from "../../components/ProductSearch";

function HomePage() {
  return (
    <div className="HomePage" data-testid="HomePage">
      <h1>Home</h1>
      <ProductSearch />
    </div>
  );
}

export default HomePage;
