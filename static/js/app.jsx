class App extends React.Component {
    render() {
        return <Home />
    }
}

class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            heroes: []
        }

        this.serverRequest = this.serverRequest.bind(this);
    }

    componentWillMount() {
        this.serverRequest();
    }

    serverRequest() {
        $.get("http://localhost:8080/api/heroes", res => {
            console.log("res... ", res);
            this.setState({
                heroes: res
            });
        });
    }

    componentDidMount() {
        this.serverRequest
    }

    render() {
        return (
            <div className="container">
                {this.state.heroes.map(function (hero, i) {
                    return <Hero hero={hero}/>
                })}
            </div>
        )
    }
}

class Hero extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div className="container" >
                <a href={`http:localhost:8080/api/hero/${this.props.hero.id}`}>{this.props.hero.name}</a>
            </div >
        )
    }
}
ReactDOM.render(<App />, document.getElementById("app"));