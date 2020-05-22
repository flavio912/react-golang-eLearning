import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Icon, { IconNames } from "sharedComponents/core/Icon";
import { Theme } from "helpers/theme";
import Button, {Archetypes} from "sharedComponents/core/Button";

const useStyles = createUseStyles((theme: Theme) => ({
	buttonRadius: {
		borderRadius: '18.5px'
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
	const archetype:Archetypes = disabled ? 'grey' : 'default';
	return (
		<Button 
			className={`${classes.buttonRadius} ${disabled ? classes.disableBtn : ''}`}
			onClick={() => onArrowClick()}
			archetype={archetype}
		>
			<Icon size={13} name={iconName} />
		</Button>
	)
}

export default PaginateButton;