import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import Blog from "./Components/Blog/Blog"
import Group from "./Media/Group.svg";
import firstwave from "./Media/firstwave.svg";


function App(): JSX.Element {
  return (
      <Router>
<>
<section>
      <img src={Group} alt="logo" className="-m-1 w-24"/>
    <div className="text-center text-white">
        <div className="text-xl">Howdy, I am</div>
        <div className="text-4xl ">Joseph Hodes</div>
        <div className="text-xl">I'm a passionate software engineer</div>
    </div>
    <img src={firstwave} alt="wave" className="absolute -bottom-20  h-48 w-full md:h-18 md:-bottom-28"/>

</section>
</>
      <Switch>
        <Route path={`/blog`}> <Blog></Blog></Route>
      </Switch>
    </Router>
  );
}

export default App;