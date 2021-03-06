import React, { Component } from "react";

import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete';


class AutocompleteSelect extends Component {
  constructor() {
    super();

    this.state = {
      options: [],
      setSelectedOption: null,
    };
  }

  componentDidMount() {
    this.setInitialOptions(this.setSelectedOption);
  }

  componentDidUpdate(prevProps) {
    if (prevProps.value === this.props.value && prevProps.triggerReload === this.props.triggerReload) {
      return;
    }

    this.setInitialOptions(this.setSelectedOption);
  }

  setInitialOptions = (callbackFunc) => {
    this.props.getOptions("", options => {
      this.setState({
        options: options,
      }, callbackFunc);
    });
  }

  setSelectedOption = () => {
    if (this.props.getOption !== undefined) {
      if (this.props.value !== undefined && this.props.value !== "" && this.props.value !== null) {
        this.props.getOption(this.props.value, resp => {
          this.setState({
            selectedOption: resp,
          });
        });
      } else {
        this.setState({
          selectedOption: "",
        });
      }
    } else {
      if (this.props.value !== undefined && this.props.value !== "" && this.props.value !== null) {
        for (const opt of this.state.options) {
          if (this.props.value === opt.value) {
            this.setState({
              selectedOption: opt,
            });
          }
        }
      } else {
        this.setState({
          selectedOption: "",
        });
      }
    }
  }

  render() {
    return(
      <Autocomplete
        id={this.props.id}
        options={this.state.options}
        getOptionLabel={(option) => option.label || ""}
        onOpen={() => this.setInitialOptions(null)}
        openOnFocus={true}
        value={this.state.selectedOption || ""}
        onChange={this.onChange}
        onInputChange={this.onInputChange}
        className={this.props.className}
        renderInput={(params) => <TextField required={!!this.props.required} placeholder={this.props.label} {...params} /> }
        disableClearable={!this.props.clearable}
      />
    );
  }

  onChange = (e, v, r) => {
    let value = null;
    if (v !== null) {
      value = v.value;
    }

    this.setState({
      selectedOption: v,
    });

    this.props.onChange({
      target: {
        id: this.props.id,
        value: value,
      },
    });
  }

  onInputChange = (e, v, r) => {
    this.props.getOptions(v, options => {
      this.setState({
        options: options,
      });
    });
  }
}

export default AutocompleteSelect;
