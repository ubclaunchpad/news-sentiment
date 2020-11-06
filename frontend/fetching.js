function getData(url) {
  const fetchedData = fetch(url)
    .then((res) => {
      if (res.ok) {
        console.log("Success");
      } else {
        console.log("Failure");
      }
      return res.json();
    })
    .then((result) => {
      console.log(result.body);
      return result;
    })
    .catch(console.error());

  return fetchedData;
}

export default getData;
