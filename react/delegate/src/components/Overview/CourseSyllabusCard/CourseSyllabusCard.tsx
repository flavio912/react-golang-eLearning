import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Icon from "sharedComponents/core/Icon";
import { Theme } from "helpers/theme";
import PaginateButton from './PaginateButton';
import classNames from "classnames";

const useStyles = createUseStyles((theme: Theme) => ({
	card: {
		display: "flex",
		flexDirection: "column",
		backgroundColor: theme.colors.primaryWhite,
		boxShadow: theme.shadows.primary,
		borderRadius: theme.primaryBorderRadius,
		border: `1px solid ${theme.colors.borderGrey}`,		
		maxWidth: '311px',
		overflow: 'hidden',
	},
	header: {
		display: 'flex',		
		lineHeight: '37px',
		letterSpacing: '-0.45px',
		fontSize: theme.fontSizes.extraLarge,
		color: theme.colors.primaryBlack,
		fontWeight: 'bold',	
		alignItems: 'center',
	},
	completeDiv: {
		width: '113px',
		height: '21px',
		lineHeight: '21px',
		borderRadius: '9.5px',
		color: theme.colors.secondaryGreen,
		backgroundColor: theme.colors.hoverGreen,		
		letterSpacing: '0.3px',
		fontSize: '12px',
		fontWeight: 'bold',
		textAlign: 'center',
		marginLeft: 'auto',
	},
	listContainer: {
		display: 'flex',
		flexDirection: 'column',
    marginTop: '7px',
    boxSizing: 'border-box',
		overflowY: 'auto',
	},
	listItem: {
		display: 'flex',
		alignItems: 'center',
		flex: '0 0 40px',
		borderBottom: `1px solid #F4F5F7`,		
		fontWeight: 'bold',
		position: 'relative',
    padding: '0',
    boxSizing: 'border-box',
		'& p': {
			margin: 0,
			padding: 0,
			fontSize: '15px',
			color: theme.colors.primaryBlack,
			letterSpacing: '-0.38px',		
			whiteSpace: 'nowrap',
    	overflow: 'hidden',
    	textOverflow: 'ellipsis',
		}
	},
	completedIcon: {
		position: 'absolute',
		right: 0,
		transform: `translateX(-100%)`
	},
	footer: {
		display: 'flex',
		backgroundColor: '#F7F9FB',
		justifyContent: 'center',
		padding: '5px 0',
		'& p': {
			display: 'flex',
			alignItems: 'center',
			justifyContent: 'center',
			width: '182px',
			height: '38px',
			borderRadius: '18.5px',
			backgroundColor: theme.colors.primaryWhite,
			border: `1px solid ${theme.colors.borderGrey}`,
			color: theme.colors.primaryBlack,
			fontSize: theme.fontSizes.default,
			fontWeight: 'bold',
			margin: '0 9.5px',
		}
  },
	

	headerSamllPading: {
		margin: '5px 10px 0 10px',
	},
	headerMediumPadding: {
		margin: '9px 20px 0 20px',
	},
	headerLargePadding: {
		margin: '20px 20px 0 30px',
	},

	listItemSamllPadding: {
		padding: '0 10px',
	},

	listItemMediumPadding: {
		padding: '0 20px',
	},

	listItemLaqrgePadding: {
		padding: '0 30px',
	},

	smallFooterPadding: {
		padding: '5px 10px',
	},
	mediumFooterPadding: {
		padding: '15px 24px',
	},
	largeFooterPadding: {
		padding: '20px 30px',
	},

}));

export type PaddingOptions = "none" | "small" | "medium" | "large";


export interface CourseSyllabus {
	completePercentage: number;
  modules: (ModulesEntity)[];
}
export interface ModulesEntity {
  sections: (SectionsEntity)[];
}
export interface SectionsEntity {
  name: string;
  uuid: string;
  complete: boolean;
}

type Props = {
	courseSyllabus: CourseSyllabus,
	padding?: PaddingOptions,
};

function CourseSyllabusCard({ courseSyllabus, padding = "large" }: Props) {

	const getMaxRows = (syllabusData:CourseSyllabus): number => {
		if (!syllabusData || syllabusData.modules.length === 0)
			return 0;
		else {
			return Math.max.apply(Math, courseSyllabus.modules.map(item => {
				return item.sections.length;
			}))
		}
	}

	const theme = useTheme();
	const classes = useStyles({ theme });
	const [curPage, setCurpage] = React.useState(0);
	const [countPerPage, setCountPerPage] = React.useState(getMaxRows(courseSyllabus));

	const paddingHeader = {
		none: "",
		small: classes.headerSamllPading,
		medium: classes.headerMediumPadding,
		large: classes.headerLargePadding,
	};

	const paddingListItem = {
		none: "",
		small: classes.listItemSamllPadding,
		medium: classes.listItemMediumPadding,
		large: classes.listItemLaqrgePadding,
	}

	const paddingFooter = {
		none: "",
		small: classes.smallFooterPadding,
		medium: classes.mediumFooterPadding,
		large: classes.largeFooterPadding,
	}	
	React.useEffect(() => {
		setCountPerPage(getMaxRows(courseSyllabus))
	}, [courseSyllabus]);

  const onClickPrev = () => {
		if (curPage === 0)
			return;
    else setCurpage(curPage-1)
	}

	const onClickNext = () => {
		if (curPage+1 >= courseSyllabus.modules.length)
			return
    else setCurpage(curPage + 1)
	}

	const renderList = () => {
		const renderComp : React.ReactElement[] = [];
		for(let i=0; i<countPerPage; i++) {			
			renderComp.push(
				<div className={classNames(classes.listItem, paddingListItem[padding])} key={i}>
					{
						(i < courseSyllabus.modules[curPage].sections.length ) &&
						<>
							<p style={{paddingRight: `${courseSyllabus.modules[curPage].sections[i].complete ? '20px' : 0}`}}>{courseSyllabus.modules[curPage].sections[i].name}</p>
							{courseSyllabus.modules[curPage].sections[i].complete && <Icon size={21} name={'CourseStatus_Completed'} className={classes.completedIcon} />}
						</>
					}
				</div>
			)
		}
		return renderComp;
  }
  
	return (
		<div className={classNames(classes.card)}>
			<div className={classNames(classes.header, paddingHeader[padding])}>
				Course Syllabus
				<div className={classes.completeDiv}>
          {`${courseSyllabus.completePercentage}% Complete`}
				</div>
			</div>
			<div className={classes.listContainer} style={{maxHeight: `${40 * countPerPage}px`}}>
				{renderList()}
			</div>
			<div className={classNames(classes.footer, paddingFooter[padding])}>
        <PaginateButton 
            iconName={curPage === 0 ? 'ArrowLeft' : 'ArrowLeftNavyBlue'}
            disabled={curPage === 0}
            onArrowClick={onClickPrev}
          />
				<p>{`Module ${courseSyllabus.modules.length > 0 ? curPage + 1 : curPage} of ${courseSyllabus.modules.length}`}</p>
				<PaginateButton 
					iconName={curPage+1 >= courseSyllabus.modules.length ? 'ArrowRight' : 'ArrowRightNavyBlue'}
					disabled={curPage+1 >= courseSyllabus.modules.length}
					onArrowClick={onClickNext}
				/>
			</div>
		</div>
	)
}

export default CourseSyllabusCard;