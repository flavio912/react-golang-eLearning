import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Icon, { IconNames } from "sharedComponents/core/Icon";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
	buttonRadius: {
		display: 'flex',
		alignItems: 'center',
		justifyContent: 'center',
		width: '40px',
		minWidth: '40px',
		height: '40px',
		borderRadius: '20px',		
		cursor: 'pointer',
		backgroundColor: 'white',	
		transition: "0.1s ease",
    	transitionProperty: "border-colour, background-color",
		outline: "none",
		border: '1px solid',
		borderColor: theme.colors.borderGrey,		
		"&:focus": {
			borderColor: theme.colors.primaryBlue,
		}
	},
	disableBtn: {
		backgroundColor: '#E9EBEB',
		cursor: "none",
		pointerEvents: "none"
	}
}));

type Props = {
	iconName: IconNames;
	disabled: boolean;
	onArrowClick: Function;
}

function PaginateButton({iconName, disabled, onArrowClick}: Props) {
	const theme = useTheme();
	const classes = useStyles({ theme });
	return (
		<button 
			className={`${classes.buttonRadius} ${disabled ? classes.disableBtn : ''}`}
			onClick={() => onArrowClick()}
		>
			<Icon size={13} name={iconName} />
		</button>
	)
}

export default PaginateButton;